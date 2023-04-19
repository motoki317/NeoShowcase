// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"fmt"
	"github.com/cert-manager/cert-manager/pkg/client/clientset/versioned"
	"github.com/fsouza/go-dockerclient"
	"github.com/google/wire"
	"github.com/traPtitech/neoshowcase/pkg/domain/web"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/admindb"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/backend/dockerimpl"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/backend/k8simpl"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/dbmanager"
	"github.com/traPtitech/neoshowcase/pkg/interface/grpc"
	"github.com/traPtitech/neoshowcase/pkg/interface/repository"
	"github.com/traPtitech/neoshowcase/pkg/usecase"
	"github.com/traefik/traefik/v2/pkg/provider/kubernetes/crd/generated/clientset/versioned/typed/traefikcontainous/v1alpha1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func NewWithDocker(c2 Config) (*Server, error) {
	config := c2.DB
	db, err := admindb.New(config)
	if err != nil {
		return nil, err
	}
	applicationRepository := repository.NewApplicationRepository(db)
	gitRepositoryRepository := repository.NewGitRepositoryRepository(db)
	publicKeys, err := provideRepositoryPublicKey(c2)
	if err != nil {
		return nil, err
	}
	buildRepository := repository.NewBuildRepository(db)
	client, err := docker.NewClientFromEnv()
	if err != nil {
		return nil, err
	}
	dockerimplConfig := c2.Docker
	backend, err := dockerimpl.NewDockerBackend(client, dockerimplConfig)
	if err != nil {
		return nil, err
	}
	logStreamService := usecase.NewLogStreamService()
	controllerBuilderService := grpc.NewControllerBuilderService(logStreamService)
	imageConfig := c2.Image
	appBuildHelper := usecase.NewAppBuildHelper(controllerBuilderService, imageConfig)
	environmentRepository := repository.NewEnvironmentRepository(db)
	controllerSSGenService := grpc.NewControllerSSGenService()
	appDeployHelper := usecase.NewAppDeployHelper(backend, applicationRepository, buildRepository, environmentRepository, controllerSSGenService, imageConfig)
	containerStateMutator := usecase.NewContainerStateMutator(applicationRepository, backend)
	continuousDeploymentService, err := usecase.NewContinuousDeploymentService(applicationRepository, buildRepository, backend, appBuildHelper, controllerBuilderService, appDeployHelper, containerStateMutator)
	if err != nil {
		return nil, err
	}
	repositoryFetcherService, err := usecase.NewRepositoryFetcherService(applicationRepository, gitRepositoryRepository, publicKeys, continuousDeploymentService)
	if err != nil {
		return nil, err
	}
	controllerServiceHandler := grpc.NewControllerService(repositoryFetcherService, continuousDeploymentService, controllerBuilderService, logStreamService)
	mainControllerServer := provideControllerServer(c2, controllerServiceHandler, controllerBuilderService, controllerSSGenService)
	artifactRepository := repository.NewArtifactRepository(db)
	storageConfig := c2.Storage
	storage, err := provideStorage(storageConfig)
	if err != nil {
		return nil, err
	}
	cleanerService, err := usecase.NewCleanerService(artifactRepository, applicationRepository, buildRepository, imageConfig, storage)
	if err != nil {
		return nil, err
	}
	server := &Server{
		controllerServer: mainControllerServer,
		db:               db,
		backend:          backend,
		cdService:        continuousDeploymentService,
		fetcherService:   repositoryFetcherService,
		cleanerService:   cleanerService,
	}
	return server, nil
}

func NewWithK8S(c2 Config) (*Server, error) {
	config := c2.DB
	db, err := admindb.New(config)
	if err != nil {
		return nil, err
	}
	applicationRepository := repository.NewApplicationRepository(db)
	gitRepositoryRepository := repository.NewGitRepositoryRepository(db)
	publicKeys, err := provideRepositoryPublicKey(c2)
	if err != nil {
		return nil, err
	}
	buildRepository := repository.NewBuildRepository(db)
	restConfig, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}
	traefikContainousV1alpha1Client, err := v1alpha1.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}
	versionedClientset, err := versioned.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}
	k8simplConfig := c2.K8s
	backend, err := k8simpl.NewK8SBackend(clientset, traefikContainousV1alpha1Client, versionedClientset, k8simplConfig)
	if err != nil {
		return nil, err
	}
	logStreamService := usecase.NewLogStreamService()
	controllerBuilderService := grpc.NewControllerBuilderService(logStreamService)
	imageConfig := c2.Image
	appBuildHelper := usecase.NewAppBuildHelper(controllerBuilderService, imageConfig)
	environmentRepository := repository.NewEnvironmentRepository(db)
	controllerSSGenService := grpc.NewControllerSSGenService()
	appDeployHelper := usecase.NewAppDeployHelper(backend, applicationRepository, buildRepository, environmentRepository, controllerSSGenService, imageConfig)
	containerStateMutator := usecase.NewContainerStateMutator(applicationRepository, backend)
	continuousDeploymentService, err := usecase.NewContinuousDeploymentService(applicationRepository, buildRepository, backend, appBuildHelper, controllerBuilderService, appDeployHelper, containerStateMutator)
	if err != nil {
		return nil, err
	}
	repositoryFetcherService, err := usecase.NewRepositoryFetcherService(applicationRepository, gitRepositoryRepository, publicKeys, continuousDeploymentService)
	if err != nil {
		return nil, err
	}
	controllerServiceHandler := grpc.NewControllerService(repositoryFetcherService, continuousDeploymentService, controllerBuilderService, logStreamService)
	mainControllerServer := provideControllerServer(c2, controllerServiceHandler, controllerBuilderService, controllerSSGenService)
	artifactRepository := repository.NewArtifactRepository(db)
	storageConfig := c2.Storage
	storage, err := provideStorage(storageConfig)
	if err != nil {
		return nil, err
	}
	cleanerService, err := usecase.NewCleanerService(artifactRepository, applicationRepository, buildRepository, imageConfig, storage)
	if err != nil {
		return nil, err
	}
	server := &Server{
		controllerServer: mainControllerServer,
		db:               db,
		backend:          backend,
		cdService:        continuousDeploymentService,
		fetcherService:   repositoryFetcherService,
		cleanerService:   cleanerService,
	}
	return server, nil
}

// wire.go:

var commonSet = wire.NewSet(web.NewServer, admindb.New, dbmanager.NewMariaDBManager, dbmanager.NewMongoDBManager, repository.NewApplicationRepository, repository.NewAvailableDomainRepository, repository.NewGitRepositoryRepository, repository.NewEnvironmentRepository, repository.NewBuildRepository, repository.NewArtifactRepository, repository.NewUserRepository, grpc.NewAPIServiceServer, grpc.NewAuthInterceptor, grpc.NewControllerService, grpc.NewControllerBuilderService, grpc.NewControllerSSGenService, usecase.NewAPIServerService, usecase.NewAppBuildHelper, usecase.NewAppDeployHelper, usecase.NewContinuousDeploymentService, usecase.NewRepositoryFetcherService, usecase.NewCleanerService, usecase.NewLogStreamService, usecase.NewContainerStateMutator, provideRepositoryPublicKey,
	provideStorage,
	provideControllerServer, wire.FieldsOf(new(Config), "DB", "Storage", "Docker", "K8s", "Image"), wire.Struct(new(Server), "*"),
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