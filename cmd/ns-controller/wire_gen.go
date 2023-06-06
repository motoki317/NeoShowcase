// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"fmt"
	"github.com/cert-manager/cert-manager/pkg/client/clientset/versioned"
	"github.com/google/wire"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/backend/dockerimpl"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/backend/k8simpl"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/dbmanager"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/grpc"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/repository"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/webhook"
	"github.com/traPtitech/neoshowcase/pkg/usecase/apiserver"
	"github.com/traPtitech/neoshowcase/pkg/usecase/cdservice"
	"github.com/traPtitech/neoshowcase/pkg/usecase/cleaner"
	"github.com/traPtitech/neoshowcase/pkg/usecase/logstream"
	"github.com/traPtitech/neoshowcase/pkg/usecase/repofetcher"
	"github.com/traPtitech/neoshowcase/pkg/usecase/sshserver"
	"github.com/traefik/traefik/v2/pkg/provider/kubernetes/crd/generated/clientset/versioned/typed/traefikio/v1alpha1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func NewWithDocker(c2 Config) (*Server, error) {
	client, err := dockerimpl.NewClientFromEnv()
	if err != nil {
		return nil, err
	}
	config := c2.Docker
	imageConfig := c2.Image
	backend, err := dockerimpl.NewDockerBackend(client, config, imageConfig)
	if err != nil {
		return nil, err
	}
	repositoryConfig := c2.DB
	db, err := repository.New(repositoryConfig)
	if err != nil {
		return nil, err
	}
	applicationRepository := repository.NewApplicationRepository(db)
	gitRepositoryRepository := repository.NewGitRepositoryRepository(db)
	publicKeys, err := providePublicKey(c2)
	if err != nil {
		return nil, err
	}
	buildRepository := repository.NewBuildRepository(db)
	service := logstream.NewService()
	controllerBuilderService := grpc.NewControllerBuilderService(service)
	environmentRepository := repository.NewEnvironmentRepository(db)
	controllerSSGenService := grpc.NewControllerSSGenService()
	appDeployHelper := cdservice.NewAppDeployHelper(backend, applicationRepository, buildRepository, environmentRepository, controllerSSGenService, imageConfig)
	containerStateMutator := cdservice.NewContainerStateMutator(applicationRepository, backend)
	cdserviceService, err := cdservice.NewService(applicationRepository, buildRepository, backend, controllerBuilderService, appDeployHelper, containerStateMutator)
	if err != nil {
		return nil, err
	}
	repofetcherService, err := repofetcher.NewService(applicationRepository, gitRepositoryRepository, publicKeys, cdserviceService)
	if err != nil {
		return nil, err
	}
	controllerServiceHandler := grpc.NewControllerService(backend, repofetcherService, cdserviceService, controllerBuilderService, service)
	mainControllerServer := provideControllerServer(c2, controllerServiceHandler, controllerBuilderService, controllerSSGenService)
	sshConfig := c2.SSH
	userRepository := repository.NewUserRepository(db)
	sshServer := sshserver.NewSSHServer(sshConfig, publicKeys, backend, applicationRepository, userRepository)
	receiverConfig := c2.Webhook
	receiver := webhook.NewReceiver(receiverConfig, gitRepositoryRepository, repofetcherService)
	artifactRepository := repository.NewArtifactRepository(db)
	storageConfig := c2.Storage
	storage, err := provideStorage(storageConfig)
	if err != nil {
		return nil, err
	}
	cleanerService, err := cleaner.NewService(artifactRepository, applicationRepository, buildRepository, imageConfig, storage)
	if err != nil {
		return nil, err
	}
	server := &Server{
		controllerServer: mainControllerServer,
		db:               db,
		backend:          backend,
		sshServer:        sshServer,
		webhook:          receiver,
		cdService:        cdserviceService,
		fetcherService:   repofetcherService,
		cleanerService:   cleanerService,
	}
	return server, nil
}

func NewWithK8S(c2 Config) (*Server, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	traefikV1alpha1Client, err := v1alpha1.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	versionedClientset, err := versioned.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	k8simplConfig := c2.K8s
	backend, err := k8simpl.NewK8SBackend(config, clientset, traefikV1alpha1Client, versionedClientset, k8simplConfig)
	if err != nil {
		return nil, err
	}
	repositoryConfig := c2.DB
	db, err := repository.New(repositoryConfig)
	if err != nil {
		return nil, err
	}
	applicationRepository := repository.NewApplicationRepository(db)
	gitRepositoryRepository := repository.NewGitRepositoryRepository(db)
	publicKeys, err := providePublicKey(c2)
	if err != nil {
		return nil, err
	}
	buildRepository := repository.NewBuildRepository(db)
	service := logstream.NewService()
	controllerBuilderService := grpc.NewControllerBuilderService(service)
	environmentRepository := repository.NewEnvironmentRepository(db)
	controllerSSGenService := grpc.NewControllerSSGenService()
	imageConfig := c2.Image
	appDeployHelper := cdservice.NewAppDeployHelper(backend, applicationRepository, buildRepository, environmentRepository, controllerSSGenService, imageConfig)
	containerStateMutator := cdservice.NewContainerStateMutator(applicationRepository, backend)
	cdserviceService, err := cdservice.NewService(applicationRepository, buildRepository, backend, controllerBuilderService, appDeployHelper, containerStateMutator)
	if err != nil {
		return nil, err
	}
	repofetcherService, err := repofetcher.NewService(applicationRepository, gitRepositoryRepository, publicKeys, cdserviceService)
	if err != nil {
		return nil, err
	}
	controllerServiceHandler := grpc.NewControllerService(backend, repofetcherService, cdserviceService, controllerBuilderService, service)
	mainControllerServer := provideControllerServer(c2, controllerServiceHandler, controllerBuilderService, controllerSSGenService)
	sshConfig := c2.SSH
	userRepository := repository.NewUserRepository(db)
	sshServer := sshserver.NewSSHServer(sshConfig, publicKeys, backend, applicationRepository, userRepository)
	receiverConfig := c2.Webhook
	receiver := webhook.NewReceiver(receiverConfig, gitRepositoryRepository, repofetcherService)
	artifactRepository := repository.NewArtifactRepository(db)
	storageConfig := c2.Storage
	storage, err := provideStorage(storageConfig)
	if err != nil {
		return nil, err
	}
	cleanerService, err := cleaner.NewService(artifactRepository, applicationRepository, buildRepository, imageConfig, storage)
	if err != nil {
		return nil, err
	}
	server := &Server{
		controllerServer: mainControllerServer,
		db:               db,
		backend:          backend,
		sshServer:        sshServer,
		webhook:          receiver,
		cdService:        cdserviceService,
		fetcherService:   repofetcherService,
		cleanerService:   cleanerService,
	}
	return server, nil
}

// wire.go:

var commonSet = wire.NewSet(dbmanager.NewMariaDBManager, dbmanager.NewMongoDBManager, repository.New, repository.NewApplicationRepository, repository.NewGitRepositoryRepository, repository.NewEnvironmentRepository, repository.NewBuildRepository, repository.NewArtifactRepository, repository.NewUserRepository, grpc.NewAPIServiceServer, grpc.NewAuthInterceptor, grpc.NewControllerService, grpc.NewControllerBuilderService, grpc.NewControllerSSGenService, webhook.NewReceiver, apiserver.NewService, cdservice.NewAppDeployHelper, cdservice.NewContainerStateMutator, cdservice.NewService, repofetcher.NewService, cleaner.NewService, logstream.NewService, sshserver.NewSSHServer, providePublicKey,
	provideStorage,
	provideControllerServer, wire.FieldsOf(new(Config), "Docker", "K8s", "SSH", "Webhook", "DB", "Storage", "Image"), wire.Struct(new(Server), "*"),
)

func New(c2 Config) (*Server, error) {
	switch c2.GetMode() {
	case ModeDocker:
		return NewWithDocker(c2)
	case ModeK8s:
		return NewWithK8S(c2)
	default:
		return nil, fmt.Errorf("unknown mode: %s", c2.Mode)
	}
}
