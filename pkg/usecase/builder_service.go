package usecase

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/docker/cli/cli/config/configfile"
	"github.com/docker/cli/cli/config/types"
	"github.com/mattn/go-shellwords"
	"github.com/moby/buildkit/session"
	"github.com/moby/buildkit/session/auth/authprovider"
	"github.com/samber/lo"

	"github.com/traPtitech/neoshowcase/pkg/infrastructure/grpc/pbconvert"

	"github.com/friendsofgo/errors"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	buildkit "github.com/moby/buildkit/client"
	"github.com/moby/buildkit/client/llb"
	"github.com/moby/buildkit/frontend/dockerfile/dockerfile2llb"
	"github.com/moby/buildkit/util/progress/progressui"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	gstatus "google.golang.org/grpc/status"

	"github.com/traPtitech/neoshowcase/pkg/domain"
	"github.com/traPtitech/neoshowcase/pkg/domain/builder"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/grpc/pb"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/repository"
	"github.com/traPtitech/neoshowcase/pkg/util/optional"
	"github.com/traPtitech/neoshowcase/pkg/util/retry"
)

const (
	buildScriptName = "neoshowcase_internal_build.sh"
)

type BuilderService interface {
	Start(ctx context.Context) error
	Shutdown(ctx context.Context) error
}

type builderService struct {
	client    domain.ControllerBuilderServiceClient
	buildkit  *buildkit.Client
	buildpack builder.BuildpackBackend
	storage   domain.Storage
	pubKey    *ssh.PublicKeys
	config    builder.ImageConfig

	artifactRepo domain.ArtifactRepository
	buildRepo    domain.BuildRepository
	gitRepo      domain.GitRepositoryRepository

	state       *state
	stateCancel func()
	statusLock  sync.Mutex
	response    chan<- *pb.BuilderResponse
	cancel      func()
}

func NewBuilderService(
	client domain.ControllerBuilderServiceClient,
	buildkit *buildkit.Client,
	buildpack builder.BuildpackBackend,
	storage domain.Storage,
	pubKey *ssh.PublicKeys,
	config builder.ImageConfig,
	artifactRepo domain.ArtifactRepository,
	buildRepo domain.BuildRepository,
	gitRepo domain.GitRepositoryRepository,
) BuilderService {
	return &builderService{
		client:       client,
		buildkit:     buildkit,
		buildpack:    buildpack,
		storage:      storage,
		pubKey:       pubKey,
		config:       config,
		artifactRepo: artifactRepo,
		buildRepo:    buildRepo,
		gitRepo:      gitRepo,
	}
}

func (s *builderService) Start(_ context.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	s.cancel = cancel

	response := make(chan *pb.BuilderResponse, 100)
	s.response = response

	go retry.Do(ctx, func(ctx context.Context) error {
		return s.client.ConnectBuilder(ctx, s.onRequest, response)
	}, 1*time.Second, 60*time.Second)
	go s.pruneLoop(ctx)

	return nil
}

func (s *builderService) Shutdown(_ context.Context) error {
	s.cancel()
	s.statusLock.Lock()
	defer s.statusLock.Unlock()
	if s.stateCancel != nil {
		s.stateCancel()
	}
	return nil
}

func (s *builderService) pruneLoop(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := s.prune(ctx)
			if err != nil {
				log.Errorf("failed to prune buildkit: %+v", err)
			}
		case <-ctx.Done():
			return
		}
	}
}

func (s *builderService) prune(ctx context.Context) error {
	return s.buildkit.Prune(ctx, nil, buildkit.PruneAll)
}

func (s *builderService) cancelBuild(buildID string) {
	s.statusLock.Lock()
	defer s.statusLock.Unlock()

	if s.state != nil && s.stateCancel != nil {
		if s.state.task.BuildID == buildID {
			s.stateCancel()
		}
	}
}

func (s *builderService) authSessions() []session.Attachable {
	if s.config.Registry.Username == "" && s.config.Registry.Password == "" {
		return nil
	}
	return []session.Attachable{authprovider.NewDockerAuthProvider(&configfile.ConfigFile{
		AuthConfigs: map[string]types.AuthConfig{
			s.config.Registry.Addr: {
				Username: s.config.Registry.Username,
				Password: s.config.Registry.Password,
			},
		},
	})}
}

func (s *builderService) onRequest(req *pb.BuilderRequest) {
	switch req.Type {
	case pb.BuilderRequest_START_BUILD:
		b := req.Body.(*pb.BuilderRequest_StartBuild).StartBuild
		err := s.tryStartTask(&builder.Task{
			ApplicationID: b.ApplicationId,
			BuildID:       b.BuildId,
			RepositoryID:  b.RepositoryId,
			Commit:        b.Commit,
			ImageName:     b.ImageName,
			ImageTag:      b.ImageTag,
			BuildConfig:   pbconvert.FromPBBuildConfig(b.BuildConfig),
		})
		if err != nil {
			log.Errorf("failed to start build: %+v", err)
		}
	case pb.BuilderRequest_CANCEL_BUILD:
		b := req.Body.(*pb.BuilderRequest_CancelBuild).CancelBuild
		s.cancelBuild(b.BuildId)
	default:
		log.Errorf("unknown builder request type: %v", req.Type)
	}
}

func (s *builderService) tryStartTask(task *builder.Task) error {
	s.statusLock.Lock()
	defer s.statusLock.Unlock()

	if s.state != nil {
		return errors.New("builder unavailable")
	}

	now := time.Now()
	err := s.buildRepo.UpdateBuild(context.Background(), task.BuildID, domain.UpdateBuildArgs{
		FromStatus: optional.From(domain.BuildStatusQueued),
		Status:     optional.From(domain.BuildStatusBuilding),
		StartedAt:  optional.From(now),
		UpdatedAt:  optional.From(now),
	})
	if err == repository.ErrNotFound {
		return nil // other builder has acquired the build lock - skip
	}
	if err != nil {
		return err
	}

	log.Infof("Starting build for %v", task.BuildID)

	repo, err := s.gitRepo.GetRepository(context.Background(), task.RepositoryID)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	finish := make(chan struct{})
	st := newState(task, repo, s.response)
	s.state = st
	s.stateCancel = func() {
		cancel()
		<-finish
	}

	go func() {
		s.response <- &pb.BuilderResponse{Type: pb.BuilderResponse_BUILD_STARTED, Body: &pb.BuilderResponse_Started{Started: &pb.BuildStarted{
			ApplicationId: task.ApplicationID,
			BuildId:       task.BuildID,
		}}}
		status := s.process(ctx, st)
		s.finalize(context.Background(), st, status) // don't want finalization tasks to be cancelled
		s.response <- &pb.BuilderResponse{Type: pb.BuilderResponse_BUILD_SETTLED, Body: &pb.BuilderResponse_Settled{Settled: &pb.BuildSettled{
			ApplicationId: task.ApplicationID,
			BuildId:       task.BuildID,
			Reason:        toPBSettleReason(status),
		}}}

		cancel()
		close(finish)
		s.statusLock.Lock()
		s.state = nil
		s.stateCancel = nil
		s.statusLock.Unlock()
		log.Infof("Build settled for %v", task.BuildID)
	}()

	return nil
}

func (s *builderService) process(ctx context.Context, st *state) domain.BuildStatus {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	go s.updateStatusLoop(ctx, st.task.BuildID)

	err := st.initTempFiles()
	if err != nil {
		log.Errorf("failed to init temp files: %+v", err)
		return domain.BuildStatusFailed
	}

	err = s.cloneRepository(ctx, st)
	if err != nil {
		log.Errorf("failed to clone repository: %+v", err)
		return domain.BuildStatusFailed
	}

	return s.build(ctx, st)
}

func (s *builderService) updateStatusLoop(ctx context.Context, buildID string) {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			err := s.buildRepo.UpdateBuild(ctx, buildID, domain.UpdateBuildArgs{UpdatedAt: optional.From(time.Now())})
			if err != nil {
				log.Errorf("failed to update build time: %+v", err)
			}
		case <-ctx.Done():
			return
		}
	}
}

func (s *builderService) cloneRepository(ctx context.Context, st *state) error {
	repo, err := git.PlainInit(st.repositoryTempDir, false)
	if err != nil {
		return errors.Wrap(err, "failed to init repository")
	}
	auth, err := domain.GitAuthMethod(st.repository, s.pubKey)
	if err != nil {
		return err
	}
	remote, err := repo.CreateRemote(&config.RemoteConfig{
		Name: "origin",
		URLs: []string{st.repository.URL},
	})
	if err != nil {
		return errors.Wrap(err, "failed to add remote")
	}
	targetRef := plumbing.NewRemoteReferenceName("origin", "target")
	err = remote.FetchContext(ctx, &git.FetchOptions{
		RemoteName: "origin",
		RefSpecs:   []config.RefSpec{config.RefSpec(fmt.Sprintf("+%s:%s", st.task.Commit, targetRef))},
		Depth:      1,
		Auth:       auth,
	})
	if err != nil {
		return errors.Wrap(err, "failed to clone repository")
	}
	wt, err := repo.Worktree()
	if err != nil {
		return errors.Wrap(err, "failed to get worktree")
	}
	err = wt.Checkout(&git.CheckoutOptions{Branch: targetRef})
	if err != nil {
		return errors.Wrap(err, "failed to checkout")
	}
	sm, err := wt.Submodules()
	if err != nil {
		return errors.Wrap(err, "getting submodules")
	}
	err = sm.Update(&git.SubmoduleUpdateOptions{
		Init:              true,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Auth:              auth,
	})
	if err != nil {
		return errors.Wrap(err, "updating submodules")
	}
	return nil
}

func (s *builderService) finalize(ctx context.Context, st *state, status domain.BuildStatus) {
	// ログファイルの保存
	if st.logTempFile != nil {
		_ = st.logTempFile.Close()
		if err := domain.SaveBuildLog(s.storage, st.logTempFile.Name(), st.task.BuildID); err != nil {
			log.Errorf("failed to save build log: %+v", err)
		}
	}

	// 生成物tarの保存
	if st.artifactTempFile != nil {
		_ = st.artifactTempFile.Close()
		if status == domain.BuildStatusSucceeded {
			err := func() error {
				filename := st.artifactTempFile.Name()
				stat, err := os.Stat(filename)
				if err != nil {
					return errors.Wrap(err, "failed to open artifact")
				}
				artifact := domain.NewArtifact(st.task.BuildID, stat.Size())
				err = s.artifactRepo.CreateArtifact(ctx, artifact)
				if err != nil {
					return errors.Wrap(err, "failed to create artifact")
				}
				err = domain.SaveArtifact(s.storage, filename, artifact.ID)
				if err != nil {
					return errors.Wrap(err, "failed to save artifact")
				}
				return nil
			}()
			if err != nil {
				log.Errorf("failed to process artifact: %+v", err)
			}
		} else {
			_ = os.Remove(st.artifactTempFile.Name())
		}
	}

	// 一時リポジトリディレクトリの削除
	if st.repositoryTempDir != "" {
		_ = os.RemoveAll(st.repositoryTempDir)
	}

	// Build更新
	now := time.Now()
	updateArgs := domain.UpdateBuildArgs{
		Status:     optional.From(status),
		UpdatedAt:  optional.From(now),
		FinishedAt: optional.From(now),
	}
	if err := s.buildRepo.UpdateBuild(ctx, st.task.BuildID, updateArgs); err != nil {
		log.Errorf("failed to update build: %+v", err)
	}
}

func (s *builderService) build(ctx context.Context, st *state) domain.BuildStatus {
	st.writeLog("[ns-builder] Build started.")

	ch := make(chan *buildkit.SolveStatus)
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() (err error) {
		switch bc := st.task.BuildConfig.(type) {
		case *domain.BuildConfigRuntimeBuildpack:
			defer close(ch)
			return s.buildImageBuildpack(ctx, st, bc)
		case *domain.BuildConfigRuntimeCmd:
			return s.buildImageWithCmd(ctx, st, ch, bc)
		case *domain.BuildConfigRuntimeDockerfile:
			return s.buildImageWithDockerfile(ctx, st, ch, bc)
		case *domain.BuildConfigStaticCmd:
			return s.buildStaticWithCmd(ctx, st, ch, bc)
		case *domain.BuildConfigStaticDockerfile:
			return s.buildStaticWithDockerfile(ctx, st, ch, bc)
		default:
			return errors.New("unknown build config type")
		}
	})
	eg.Go(func() error {
		// ビルドログを収集
		// TODO: VertexWarningを使う (LLBのどのvertexに問題があったか)
		_, err := progressui.DisplaySolveStatus(ctx, "", nil, st.getLogWriter(), ch)
		return err
	})

	err := eg.Wait()
	if err != nil {
		if err == context.Canceled || err == context.DeadlineExceeded || errors.Is(err, gstatus.FromContextError(context.Canceled).Err()) {
			st.writeLog("[ns-builder] Build cancelled.")
			return domain.BuildStatusCanceled
		}
		log.Errorf("failed to build: %+v", err)
		return domain.BuildStatusFailed
	}

	st.writeLog("[ns-builder] Build succeeded!")
	return domain.BuildStatusSucceeded
}

func (s *builderService) buildImageBuildpack(
	ctx context.Context,
	st *state,
	bc *domain.BuildConfigRuntimeBuildpack,
) error {
	contextDir := lo.Ternary(bc.Context != "", bc.Context, ".")
	buildDir := filepath.Join(st.repositoryTempDir, contextDir)
	return s.buildpack.Pack(ctx, buildDir, st.getLogWriter(), st.task.DestImage())
}

func (s *builderService) buildImageWithCmd(
	ctx context.Context,
	st *state,
	ch chan *buildkit.SolveStatus,
	bc *domain.BuildConfigRuntimeCmd,
) error {
	var ls llb.State
	if bc.BaseImage == "" {
		ls = llb.Scratch()
	} else {
		ls = llb.Image(bc.BaseImage)
	}
	ls = ls.
		Dir("/srv").
		File(llb.Copy(llb.Local("local-src"), ".", ".", &llb.CopyInfo{
			CopyDirContentsOnly: true,
			AllowWildcard:       true,
			CreateDestPath:      true,
		}))

	if bc.BuildCmd != "" {
		if bc.BuildCmdShell {
			err := createScriptFile(filepath.Join(st.repositoryTempDir, buildScriptName), bc.BuildCmd)
			if err != nil {
				return err
			}
			ls = ls.Run(llb.Args([]string{"./" + buildScriptName})).Root()
		} else {
			args, err := shellwords.Parse(bc.BuildCmd)
			if err != nil {
				return err
			}
			ls = ls.Run(llb.Args(args)).Root()
		}
	}

	def, err := ls.Marshal(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to marshal llb")
	}

	_, err = s.buildkit.Solve(ctx, def, buildkit.SolveOpt{
		Exports: []buildkit.ExportEntry{{
			Type: buildkit.ExporterImage,
			Attrs: map[string]string{
				"name": st.task.DestImage(),
				"push": "true",
			},
		}},
		LocalDirs: map[string]string{
			"local-src": st.repositoryTempDir,
		},
		Session: s.authSessions(),
	}, ch)
	return err
}

func (s *builderService) buildImageWithDockerfile(
	ctx context.Context,
	st *state,
	ch chan *buildkit.SolveStatus,
	bc *domain.BuildConfigRuntimeDockerfile,
) error {
	contextDir := lo.Ternary(bc.Context != "", bc.Context, ".")
	dockerfileDir := filepath.Join(contextDir, filepath.Dir(bc.DockerfileName))
	_, err := s.buildkit.Solve(ctx, nil, buildkit.SolveOpt{
		Exports: []buildkit.ExportEntry{{
			Type: buildkit.ExporterImage,
			Attrs: map[string]string{
				"name": st.task.DestImage(),
				"push": "true",
			},
		}},
		LocalDirs: map[string]string{
			"context":    filepath.Join(st.repositoryTempDir, contextDir),
			"dockerfile": filepath.Join(st.repositoryTempDir, dockerfileDir),
		},
		Frontend:      "dockerfile.v0",
		FrontendAttrs: map[string]string{"filename": bc.DockerfileName},
		Session:       s.authSessions(),
	}, ch)
	return err
}

func (s *builderService) buildStaticWithCmd(
	ctx context.Context,
	st *state,
	ch chan *buildkit.SolveStatus,
	bc *domain.BuildConfigStaticCmd,
) error {
	var ls llb.State
	if bc.BaseImage == "" {
		ls = llb.Scratch()
	} else {
		ls = llb.Image(bc.BaseImage)
	}
	ls = ls.
		Dir("/srv").
		File(llb.Copy(llb.Local("local-src"), ".", ".", &llb.CopyInfo{
			CopyDirContentsOnly: true,
			AllowWildcard:       true,
			CreateDestPath:      true,
		}))

	if bc.BuildCmd != "" {
		if bc.BuildCmdShell {
			err := createScriptFile(filepath.Join(st.repositoryTempDir, buildScriptName), bc.BuildCmd)
			if err != nil {
				return err
			}
			ls = ls.Run(llb.Args([]string{"./" + buildScriptName})).Root()
		} else {
			args, err := shellwords.Parse(bc.BuildCmd)
			if err != nil {
				return err
			}
			ls = ls.Run(llb.Args(args)).Root()
		}
	}

	// ビルドで生成された静的ファイルのみを含むScratchイメージを構成
	def, err := llb.
		Scratch().
		File(llb.Copy(ls, filepath.Join("/srv", bc.ArtifactPath), "/", &llb.CopyInfo{
			CopyDirContentsOnly: true,
			CreateDestPath:      true,
			AllowWildcard:       true,
		})).
		Marshal(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to marshal llb")
	}

	_, err = s.buildkit.Solve(ctx, def, buildkit.SolveOpt{
		Exports: []buildkit.ExportEntry{{
			Type:   buildkit.ExporterTar,
			Output: func(_ map[string]string) (io.WriteCloser, error) { return st.artifactTempFile, nil },
		}},
		LocalDirs: map[string]string{
			"local-src": st.repositoryTempDir,
		},
		Session: s.authSessions(),
	}, ch)
	return err
}

func (s *builderService) buildStaticWithDockerfile(
	ctx context.Context,
	st *state,
	ch chan *buildkit.SolveStatus,
	bc *domain.BuildConfigStaticDockerfile,
) error {
	contextDir := lo.Ternary(bc.Context != "", bc.Context, filepath.Dir(bc.DockerfileName))
	dockerfile, err := os.ReadFile(filepath.Join(st.repositoryTempDir, contextDir, bc.DockerfileName))
	if err != nil {
		return err
	}

	b, _, _, err := dockerfile2llb.Dockerfile2LLB(ctx, dockerfile, dockerfile2llb.ConvertOpt{})
	if err != nil {
		return err
	}

	def, err := llb.
		Scratch().
		File(llb.Copy(*b, bc.ArtifactPath, "/", &llb.CopyInfo{
			CopyDirContentsOnly: true,
			CreateDestPath:      true,
			AllowWildcard:       true,
		})).
		Marshal(ctx)
	if err != nil {
		return err
	}

	_, err = s.buildkit.Solve(ctx, def, buildkit.SolveOpt{
		Exports: []buildkit.ExportEntry{{
			Type:   buildkit.ExporterTar,
			Output: func(_ map[string]string) (io.WriteCloser, error) { return st.artifactTempFile, nil },
		}},
		LocalDirs: map[string]string{
			"context": filepath.Join(st.repositoryTempDir, contextDir),
		},
		Session: s.authSessions(),
	}, ch)
	return err
}

func createScriptFile(filename string, script string) error {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString("#!/bin/sh\nset -eux\n" + script)
	if err != nil {
		return err
	}
	return nil
}

type state struct {
	task       *builder.Task
	repository *domain.Repository
	response   chan<- *pb.BuilderResponse

	repositoryTempDir string
	logTempFile       *os.File
	logWriter         *logWriter
	artifactTempFile  *os.File
}

type logWriter struct {
	buildID  string
	response chan<- *pb.BuilderResponse
	logFile  *os.File
}

func (l *logWriter) toBuilderResponse(p []byte) *pb.BuilderResponse {
	return &pb.BuilderResponse{Type: pb.BuilderResponse_BUILD_LOG, Body: &pb.BuilderResponse_Log{
		Log: &pb.BuildLogPortion{BuildId: l.buildID, Log: p},
	}}
}

func (l *logWriter) Write(p []byte) (n int, err error) {
	n, err = l.logFile.Write(p)
	if err != nil {
		return
	}
	select {
	case l.response <- l.toBuilderResponse(p):
	default:
	}
	return
}

func newState(task *builder.Task, repo *domain.Repository, response chan<- *pb.BuilderResponse) *state {
	return &state{
		task:       task,
		repository: repo,
		response:   response,
	}
}

func (s *state) static() bool {
	return s.task.BuildConfig.BuildType().DeployType() == domain.DeployTypeStatic
}

func (s *state) initTempFiles() error {
	var err error

	// ログ用一時ファイル作成
	s.logTempFile, err = os.CreateTemp("", "buildlog-")
	if err != nil {
		return errors.Wrap(err, "failed to create tmp log file")
	}
	s.logWriter = &logWriter{
		buildID:  s.task.BuildID,
		response: s.response,
		logFile:  s.logTempFile,
	}

	// 成果物tarの一時保存先作成
	if s.static() {
		s.artifactTempFile, err = os.CreateTemp("", "artifacts-")
		if err != nil {
			return errors.Wrap(err, "failed to create tmp artifact file")
		}
	}

	// リポジトリクローン用の一時ディレクトリ作成
	s.repositoryTempDir, err = os.MkdirTemp("", "repository-")
	if err != nil {
		return errors.Wrap(err, "failed to create tmp repository dir")
	}

	return nil
}

func (s *state) getLogWriter() io.Writer {
	return s.logWriter
}

func (s *state) writeLog(a ...interface{}) {
	_, _ = fmt.Fprintln(s.logWriter, a...)
}

func toPBSettleReason(status domain.BuildStatus) pb.BuildSettled_Reason {
	switch status {
	case domain.BuildStatusSucceeded:
		return pb.BuildSettled_SUCCESS
	case domain.BuildStatusFailed:
		return pb.BuildSettled_FAILED
	case domain.BuildStatusCanceled:
		return pb.BuildSettled_CANCELLED
	default:
		panic(fmt.Sprintf("unexpected settled status: %v", status))
	}
}
