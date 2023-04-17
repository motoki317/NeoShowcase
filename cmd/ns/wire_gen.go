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
	"github.com/leandro-lugaresi/hub"
	"github.com/traPtitech/neoshowcase/pkg/domain/web"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/admindb"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/backend/dockerimpl"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/backend/k8simpl"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/dbmanager"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/eventbus"
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
	hubHub := hub.New()
	bus := eventbus.NewLocal(hubHub)
	config := c2.DB
	db, err := admindb.New(config)
	if err != nil {
		return nil, err
	}
	artifactRepository := repository.NewArtifactRepository(db)
	applicationRepository := repository.NewApplicationRepository(db)
	availableDomainRepository := repository.NewAvailableDomainRepository(db)
	buildRepository := repository.NewBuildRepository(db)
	environmentRepository := repository.NewEnvironmentRepository(db)
	gitRepositoryRepository := repository.NewGitRepositoryRepository(db)
	storageConfig := c2.Storage
	storage, err := provideStorage(storageConfig)
	if err != nil {
		return nil, err
	}
	logStreamService := usecase.NewLogStreamService()
	componentService := grpc.NewComponentServiceServer(bus, logStreamService)
	mariaDBConfig := c2.MariaDB
	mariaDBManager, err := dbmanager.NewMariaDBManager(mariaDBConfig)
	if err != nil {
		return nil, err
	}
	mongoDBConfig := c2.MongoDB
	mongoDBManager, err := dbmanager.NewMongoDBManager(mongoDBConfig)
	if err != nil {
		return nil, err
	}
	containerLogger, err := provideContainerLogger(c2)
	if err != nil {
		return nil, err
	}
	apiServerService := usecase.NewAPIServerService(bus, artifactRepository, applicationRepository, availableDomainRepository, buildRepository, environmentRepository, gitRepositoryRepository, storage, componentService, mariaDBManager, mongoDBManager, containerLogger, logStreamService)
	publicKeys, err := provideRepositoryPublicKey(c2)
	if err != nil {
		return nil, err
	}
	apiServiceHandler := grpc.NewAPIServiceServer(apiServerService, publicKeys)
	userRepository := repository.NewUserRepository(db)
	authInterceptor := grpc.NewAuthInterceptor(userRepository)
	mainWebAppServer := provideWebAppServer(c2, apiServiceHandler, authInterceptor)
	mainWebComponentServer := provideWebComponentServer(c2, componentService)
	client, err := docker.NewClientFromEnv()
	if err != nil {
		return nil, err
	}
	dockerimplConfig := c2.Docker
	backend, err := dockerimpl.NewDockerBackend(client, bus, dockerimplConfig)
	if err != nil {
		return nil, err
	}
	imageConfig := c2.Image
	appBuildHelper := usecase.NewAppBuildHelper(componentService, imageConfig)
	appDeployHelper := usecase.NewAppDeployHelper(backend, applicationRepository, buildRepository, environmentRepository, componentService, imageConfig)
	containerStateMutator := usecase.NewContainerStateMutator(bus, applicationRepository, backend)
	continuousDeploymentService, err := usecase.NewContinuousDeploymentService(bus, applicationRepository, buildRepository, backend, appBuildHelper, appDeployHelper, containerStateMutator)
	if err != nil {
		return nil, err
	}
	repositoryFetcherService, err := usecase.NewRepositoryFetcherService(bus, applicationRepository, gitRepositoryRepository, publicKeys)
	if err != nil {
		return nil, err
	}
	cleanerService, err := usecase.NewCleanerService(artifactRepository, applicationRepository, buildRepository, imageConfig, storage)
	if err != nil {
		return nil, err
	}
	server := &Server{
		appServer:       mainWebAppServer,
		componentServer: mainWebComponentServer,
		db:              db,
		backend:         backend,
		bus:             bus,
		cdService:       continuousDeploymentService,
		fetcherService:  repositoryFetcherService,
		cleanerService:  cleanerService,
	}
	return server, nil
}

func NewWithK8S(c2 Config) (*Server, error) {
	hubHub := hub.New()
	bus := eventbus.NewLocal(hubHub)
	config := c2.DB
	db, err := admindb.New(config)
	if err != nil {
		return nil, err
	}
	artifactRepository := repository.NewArtifactRepository(db)
	applicationRepository := repository.NewApplicationRepository(db)
	availableDomainRepository := repository.NewAvailableDomainRepository(db)
	buildRepository := repository.NewBuildRepository(db)
	environmentRepository := repository.NewEnvironmentRepository(db)
	gitRepositoryRepository := repository.NewGitRepositoryRepository(db)
	storageConfig := c2.Storage
	storage, err := provideStorage(storageConfig)
	if err != nil {
		return nil, err
	}
	logStreamService := usecase.NewLogStreamService()
	componentService := grpc.NewComponentServiceServer(bus, logStreamService)
	mariaDBConfig := c2.MariaDB
	mariaDBManager, err := dbmanager.NewMariaDBManager(mariaDBConfig)
	if err != nil {
		return nil, err
	}
	mongoDBConfig := c2.MongoDB
	mongoDBManager, err := dbmanager.NewMongoDBManager(mongoDBConfig)
	if err != nil {
		return nil, err
	}
	containerLogger, err := provideContainerLogger(c2)
	if err != nil {
		return nil, err
	}
	apiServerService := usecase.NewAPIServerService(bus, artifactRepository, applicationRepository, availableDomainRepository, buildRepository, environmentRepository, gitRepositoryRepository, storage, componentService, mariaDBManager, mongoDBManager, containerLogger, logStreamService)
	publicKeys, err := provideRepositoryPublicKey(c2)
	if err != nil {
		return nil, err
	}
	apiServiceHandler := grpc.NewAPIServiceServer(apiServerService, publicKeys)
	userRepository := repository.NewUserRepository(db)
	authInterceptor := grpc.NewAuthInterceptor(userRepository)
	mainWebAppServer := provideWebAppServer(c2, apiServiceHandler, authInterceptor)
	mainWebComponentServer := provideWebComponentServer(c2, componentService)
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
	backend, err := k8simpl.NewK8SBackend(bus, clientset, traefikContainousV1alpha1Client, versionedClientset, k8simplConfig)
	if err != nil {
		return nil, err
	}
	imageConfig := c2.Image
	appBuildHelper := usecase.NewAppBuildHelper(componentService, imageConfig)
	appDeployHelper := usecase.NewAppDeployHelper(backend, applicationRepository, buildRepository, environmentRepository, componentService, imageConfig)
	containerStateMutator := usecase.NewContainerStateMutator(bus, applicationRepository, backend)
	continuousDeploymentService, err := usecase.NewContinuousDeploymentService(bus, applicationRepository, buildRepository, backend, appBuildHelper, appDeployHelper, containerStateMutator)
	if err != nil {
		return nil, err
	}
	repositoryFetcherService, err := usecase.NewRepositoryFetcherService(bus, applicationRepository, gitRepositoryRepository, publicKeys)
	if err != nil {
		return nil, err
	}
	cleanerService, err := usecase.NewCleanerService(artifactRepository, applicationRepository, buildRepository, imageConfig, storage)
	if err != nil {
		return nil, err
	}
	server := &Server{
		appServer:       mainWebAppServer,
		componentServer: mainWebComponentServer,
		db:              db,
		backend:         backend,
		bus:             bus,
		cdService:       continuousDeploymentService,
		fetcherService:  repositoryFetcherService,
		cleanerService:  cleanerService,
	}
	return server, nil
}

// wire.go:

var commonSet = wire.NewSet(web.NewServer, hub.New, eventbus.NewLocal, admindb.New, dbmanager.NewMariaDBManager, dbmanager.NewMongoDBManager, repository.NewApplicationRepository, repository.NewAvailableDomainRepository, repository.NewGitRepositoryRepository, repository.NewEnvironmentRepository, repository.NewBuildRepository, repository.NewArtifactRepository, repository.NewUserRepository, grpc.NewAPIServiceServer, grpc.NewAuthInterceptor, grpc.NewComponentServiceServer, usecase.NewAPIServerService, usecase.NewAppBuildHelper, usecase.NewAppDeployHelper, usecase.NewContinuousDeploymentService, usecase.NewRepositoryFetcherService, usecase.NewCleanerService, usecase.NewLogStreamService, usecase.NewContainerStateMutator, provideRepositoryPublicKey,
	provideStorage,
	provideContainerLogger,
	provideWebAppServer,
	provideWebComponentServer, wire.FieldsOf(new(Config), "DB", "MariaDB", "MongoDB", "Storage", "Docker", "K8s", "Image"), wire.Struct(new(Server), "*"),
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
