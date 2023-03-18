package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/traPtitech/neoshowcase/pkg/domain"
	"github.com/traPtitech/neoshowcase/pkg/domain/builder"
	"github.com/traPtitech/neoshowcase/pkg/interface/grpc/pb"
	mock_pb "github.com/traPtitech/neoshowcase/pkg/interface/grpc/pb/mock"
	mock_repository "github.com/traPtitech/neoshowcase/pkg/interface/repository/mock"
)

func TestAppBuildService_QueueBuild(t *testing.T) {
	t.Parallel()

	t.Run("ビルドキューへの追加(Image)", func(t *testing.T) {
		t.Parallel()
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		appRepo := mock_repository.NewMockApplicationRepository(mockCtrl)
		buildRepo := mock_repository.NewMockBuildRepository(mockCtrl)
		c := mock_pb.NewMockBuilderServiceClient(mockCtrl)
		s := NewAppBuildService(appRepo, buildRepo, c, "TestRegistry", "TestPrefix")
		app := &domain.Application{
			ID: "bee2466e-9d46-45e5-a6c4-4d359504c10c",
			Repository: domain.Repository{
				URL: "https://git.trap.jp/hijiki51/git-test",
			},
			BranchName: "main",
			BuildType:  builder.BuildTypeRuntime,
		}
		build := &domain.Build{
			ID:            "f01691dd-985a-48c9-8b47-205af468431a",
			Commit:        "e46b36b48c2cc692c558502e1f57466d93bc031d",
			Status:        builder.BuildStatusQueued,
			ApplicationID: app.ID,
		}

		appRepo.EXPECT().
			GetApplication(context.Background(), app.ID).Return(app, nil)

		buildRepo.EXPECT().
			CreateBuild(context.Background(), app.ID, build.Commit).Return(build, nil)

		c.EXPECT().
			GetStatus(context.Background(), &emptypb.Empty{}).
			Return(&pb.GetStatusResponse{
				Status:  pb.BuilderStatus_WAITING,
				BuildId: build.ID,
			}, nil).
			AnyTimes()

		c.EXPECT().
			StartBuildImage(context.Background(), &pb.StartBuildImageRequest{
				ImageName: "TestRegistry/TestPrefixbee2466e-9d46-45e5-a6c4-4d359504c10c",
				ImageTag:  build.ID,
				Source: &pb.BuildSource{
					RepositoryUrl: app.Repository.URL,
					Commit:        build.Commit,
				},
				Options:       &pb.BuildOptions{},
				BuildId:       build.ID,
				ApplicationId: app.ID,
			}).
			Return(&pb.StartBuildImageResponse{}, nil)

		_, err := s.QueueBuild(context.Background(), app, build.Commit)
		s.Shutdown()
		require.Nil(t, err)
	})

	t.Run("ビルドキューへの追加(Static)", func(t *testing.T) {
		t.Parallel()
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		appRepo := mock_repository.NewMockApplicationRepository(mockCtrl)
		buildRepo := mock_repository.NewMockBuildRepository(mockCtrl)
		c := mock_pb.NewMockBuilderServiceClient(mockCtrl)
		s := NewAppBuildService(appRepo, buildRepo, c, "TestRegistry", "TestPrefix")
		app := &domain.Application{
			ID: "d563e2de-7905-4267-8a9c-51520aac02b3",
			Repository: domain.Repository{
				URL: "https://git.trap.jp/hijiki51/git-test",
			},
			BranchName: "develop",
			BuildType:  builder.BuildTypeStatic,
		}
		build := &domain.Build{
			ID:            "f01691dd-985a-48c9-8b47-205af468431a",
			Commit:        "e46b36b48c2cc692c558502e1f57466d93bc031d",
			Status:        builder.BuildStatusQueued,
			ApplicationID: app.ID,
		}

		appRepo.EXPECT().
			GetApplication(context.Background(), app.ID).Return(app, nil)

		buildRepo.EXPECT().
			CreateBuild(context.Background(), app.ID, build.Commit).Return(build, nil)

		c.EXPECT().
			GetStatus(context.Background(), &emptypb.Empty{}).
			Return(&pb.GetStatusResponse{
				Status:  pb.BuilderStatus_WAITING,
				BuildId: build.ID,
			}, nil).
			AnyTimes()

		c.EXPECT().
			StartBuildStatic(context.Background(), &pb.StartBuildStaticRequest{
				Source: &pb.BuildSource{
					RepositoryUrl: app.Repository.URL,
					Commit:        build.Commit,
				},
				Options:       &pb.BuildOptions{},
				BuildId:       build.ID,
				ApplicationId: app.ID,
			}).
			Return(&pb.StartBuildStaticResponse{}, nil)

		_, err := s.QueueBuild(context.Background(), app, build.Commit)
		s.Shutdown()
		require.Nil(t, err)
	})
	t.Run("追加されたジョブのキャンセル", func(t *testing.T) {
		t.Parallel()
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		appRepo := mock_repository.NewMockApplicationRepository(mockCtrl)
		buildLog := mock_repository.NewMockBuildRepository(mockCtrl)
		c := mock_pb.NewMockBuilderServiceClient(mockCtrl)
		s := NewAppBuildService(appRepo, buildLog, c, "TestRegistry", "TestPrefix")
		queue := s.(*appBuildService).queue

		app1 := &domain.Application{
			ID: "d563e2de-7905-4267-8a9c-51520aac02b3",
			Repository: domain.Repository{
				URL: "https://git.trap.jp/hijiki51/git-test",
			},
			BranchName: "develop",
			BuildType:  builder.BuildTypeStatic,
		}
		app2 := &domain.Application{
			ID: "19005490-5119-40ef-95e2-24a193e64a38",
			Repository: domain.Repository{
				URL: "https://git.trap.jp/hijiki51/git-test",
			},
			BranchName: "main",
			BuildType:  builder.BuildTypeStatic,
		}
		build1 := &domain.Build{
			ID:            "f01691dd-985a-48c9-8b47-205af468431a",
			Commit:        "e46b36b48c2cc692c558502e1f57466d93bc031d",
			Status:        builder.BuildStatusQueued,
			ApplicationID: app1.ID,
		}
		build2 := &domain.Build{
			ID:            "4bd30598-2962-416a-86b5-635899a96a65",
			Commit:        "fc703e1553578709c3d98e8b6468e8abbde77b54",
			Status:        builder.BuildStatusQueued,
			ApplicationID: app2.ID,
		}

		appRepo.EXPECT().
			GetApplication(context.Background(), app1.ID).Return(app1, nil)
		appRepo.EXPECT().
			GetApplication(context.Background(), app2.ID).Return(app2, nil)

		buildLog.EXPECT().CreateBuild(context.Background(), app1.ID, build1.Commit).Return(build1, nil)
		buildLog.EXPECT().CreateBuild(context.Background(), app2.ID, build2.Commit).Return(build2, nil)

		// stop processing queue
		c.EXPECT().
			GetStatus(context.Background(), &emptypb.Empty{}).
			Return(&pb.GetStatusResponse{
				Status:  pb.BuilderStatus_UNAVAILABLE,
				BuildId: "",
			}, nil).
			AnyTimes()

		id1, err := s.QueueBuild(context.Background(), app1, build1.Commit)
		if err != nil {
			t.Fatal(err)
		}
		id2, err := s.QueueBuild(context.Background(), app2, build2.Commit)
		if err != nil {
			t.Fatal(err)
		}

		// wait for the Pop() of first item
		time.Sleep(queueCheckInterval * 2)

		require.Equal(t, queue.Len(), 1)

		// could not cancel the latest one for now
		err = s.CancelBuild(context.Background(), id1)
		require.Equal(t, queue.Len(), 1)
		require.NotNil(t, err)

		// cancel waiting job
		err = s.CancelBuild(context.Background(), id2)
		require.Equal(t, queue.Len(), 0)
		require.Nil(t, err)
	})
}
