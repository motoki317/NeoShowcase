// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"fmt"
	"github.com/fsouza/go-dockerclient"
	"github.com/google/wire"
	"github.com/leandro-lugaresi/hub"
	"github.com/traPtitech/neoshowcase/pkg/domain/web"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/admindb"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/backend/dockerimpl"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/backend/k8simpl"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/dbmanager"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/eventbus"
	"github.com/traPtitech/neoshowcase/pkg/interface/broker"
	"github.com/traPtitech/neoshowcase/pkg/interface/grpc"
	"github.com/traPtitech/neoshowcase/pkg/interface/repository"
	"github.com/traPtitech/neoshowcase/pkg/usecase"
	"github.com/traefik/traefik/v2/pkg/provider/kubernetes/crd/generated/clientset/versioned/typed/traefik/v1alpha1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func NewWithDocker(c2 Config) (*Server, error) {
	server := grpc.NewServer()
	tcpListenPort := provideGRPCPort(c2)
	hubHub := hub.New()
	bus := eventbus.NewLocal(hubHub)
	config := c2.DB
	db, err := admindb.New(config)
	if err != nil {
		return nil, err
	}
	applicationRepository := repository.NewApplicationRepository(db)
	availableDomainRepository := repository.NewAvailableDomainRepository(db)
	buildRepository := repository.NewBuildRepository(db)
	environmentRepository := repository.NewEnvironmentRepository(db)
	gitRepositoryRepository := repository.NewGitRepositoryRepository(db)
	client, err := docker.NewClientFromEnv()
	if err != nil {
		return nil, err
	}
	ingressConfDirPath := _wireIngressConfDirPathValue
	backend, err := dockerimpl.NewDockerBackend(client, bus, ingressConfDirPath)
	if err != nil {
		return nil, err
	}
	staticSiteServiceClientConfig := c2.SSGen
	staticSiteServiceClientConn, err := grpc.NewStaticSiteServiceClientConn(staticSiteServiceClientConfig)
	if err != nil {
		return nil, err
	}
	staticSiteServiceClient := grpc.NewStaticSiteServiceClient(staticSiteServiceClientConn)
	dockerImageRegistryString := provideImageRegistry(c2)
	dockerImageNamePrefixString := provideImagePrefix(c2)
	appDeployService := usecase.NewAppDeployService(bus, backend, applicationRepository, buildRepository, environmentRepository, staticSiteServiceClient, dockerImageRegistryString, dockerImageNamePrefixString)
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
	apiServerService := usecase.NewAPIServerService(bus, applicationRepository, availableDomainRepository, buildRepository, environmentRepository, gitRepositoryRepository, appDeployService, backend, mariaDBManager, mongoDBManager)
	applicationService := grpc.NewApplicationServiceServer(apiServerService, availableDomainRepository)
	router := &Router{}
	webConfig := provideWebServerConfig(router)
	webServer := web.NewServer(webConfig)
	builderServiceClientConfig := c2.Builder
	builderServiceClientConn, err := grpc.NewBuilderServiceClientConn(builderServiceClientConfig)
	if err != nil {
		return nil, err
	}
	builderServiceClient := grpc.NewBuilderServiceClient(builderServiceClientConn)
	builderEventsBroker, err := broker.NewBuilderEventsBroker(builderServiceClient, bus)
	if err != nil {
		return nil, err
	}
	appBuildService := usecase.NewAppBuildService(applicationRepository, buildRepository, builderServiceClient, dockerImageRegistryString, dockerImageNamePrefixString)
	continuousDeploymentService := usecase.NewContinuousDeploymentService(bus, applicationRepository, buildRepository, environmentRepository, appDeployService, appBuildService)
	repositoryFetcherCacheDir := provideRepositoryFetcherCacheDir(c2)
	repositoryFetcherService, err := usecase.NewRepositoryFetcherService(bus, applicationRepository, repositoryFetcherCacheDir)
	if err != nil {
		return nil, err
	}
	mainServer := &Server{
		grpcServer:          server,
		grpcPort:            tcpListenPort,
		appService:          applicationService,
		webserver:           webServer,
		db:                  db,
		builderConn:         builderServiceClientConn,
		ssgenConn:           staticSiteServiceClientConn,
		backend:             backend,
		bus:                 bus,
		builderEventsBroker: builderEventsBroker,
		cdService:           continuousDeploymentService,
		fetcherService:      repositoryFetcherService,
	}
	return mainServer, nil
}

var (
	_wireIngressConfDirPathValue = dockerimpl.IngressConfDirPath("/opt/traefik/conf")
)

func NewWithK8S(c2 Config) (*Server, error) {
	server := grpc.NewServer()
	tcpListenPort := provideGRPCPort(c2)
	hubHub := hub.New()
	bus := eventbus.NewLocal(hubHub)
	config := c2.DB
	db, err := admindb.New(config)
	if err != nil {
		return nil, err
	}
	applicationRepository := repository.NewApplicationRepository(db)
	availableDomainRepository := repository.NewAvailableDomainRepository(db)
	buildRepository := repository.NewBuildRepository(db)
	environmentRepository := repository.NewEnvironmentRepository(db)
	gitRepositoryRepository := repository.NewGitRepositoryRepository(db)
	restConfig, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}
	traefikV1alpha1Client, err := v1alpha1.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}
	backend, err := k8simpl.NewK8SBackend(bus, clientset, traefikV1alpha1Client)
	if err != nil {
		return nil, err
	}
	staticSiteServiceClientConfig := c2.SSGen
	staticSiteServiceClientConn, err := grpc.NewStaticSiteServiceClientConn(staticSiteServiceClientConfig)
	if err != nil {
		return nil, err
	}
	staticSiteServiceClient := grpc.NewStaticSiteServiceClient(staticSiteServiceClientConn)
	dockerImageRegistryString := provideImageRegistry(c2)
	dockerImageNamePrefixString := provideImagePrefix(c2)
	appDeployService := usecase.NewAppDeployService(bus, backend, applicationRepository, buildRepository, environmentRepository, staticSiteServiceClient, dockerImageRegistryString, dockerImageNamePrefixString)
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
	apiServerService := usecase.NewAPIServerService(bus, applicationRepository, availableDomainRepository, buildRepository, environmentRepository, gitRepositoryRepository, appDeployService, backend, mariaDBManager, mongoDBManager)
	applicationService := grpc.NewApplicationServiceServer(apiServerService, availableDomainRepository)
	router := &Router{}
	webConfig := provideWebServerConfig(router)
	webServer := web.NewServer(webConfig)
	builderServiceClientConfig := c2.Builder
	builderServiceClientConn, err := grpc.NewBuilderServiceClientConn(builderServiceClientConfig)
	if err != nil {
		return nil, err
	}
	builderServiceClient := grpc.NewBuilderServiceClient(builderServiceClientConn)
	builderEventsBroker, err := broker.NewBuilderEventsBroker(builderServiceClient, bus)
	if err != nil {
		return nil, err
	}
	appBuildService := usecase.NewAppBuildService(applicationRepository, buildRepository, builderServiceClient, dockerImageRegistryString, dockerImageNamePrefixString)
	continuousDeploymentService := usecase.NewContinuousDeploymentService(bus, applicationRepository, buildRepository, environmentRepository, appDeployService, appBuildService)
	repositoryFetcherCacheDir := provideRepositoryFetcherCacheDir(c2)
	repositoryFetcherService, err := usecase.NewRepositoryFetcherService(bus, applicationRepository, repositoryFetcherCacheDir)
	if err != nil {
		return nil, err
	}
	mainServer := &Server{
		grpcServer:          server,
		grpcPort:            tcpListenPort,
		appService:          applicationService,
		webserver:           webServer,
		db:                  db,
		builderConn:         builderServiceClientConn,
		ssgenConn:           staticSiteServiceClientConn,
		backend:             backend,
		bus:                 bus,
		builderEventsBroker: builderEventsBroker,
		cdService:           continuousDeploymentService,
		fetcherService:      repositoryFetcherService,
	}
	return mainServer, nil
}

// wire.go:

var commonSet = wire.NewSet(web.NewServer, hub.New, eventbus.NewLocal, admindb.New, dbmanager.NewMariaDBManager, dbmanager.NewMongoDBManager, repository.NewApplicationRepository, repository.NewAvailableDomainRepository, repository.NewGitRepositoryRepository, repository.NewEnvironmentRepository, repository.NewBuildRepository, grpc.NewServer, grpc.NewApplicationServiceServer, grpc.NewBuilderServiceClientConn, grpc.NewStaticSiteServiceClientConn, grpc.NewBuilderServiceClient, grpc.NewStaticSiteServiceClient, broker.NewBuilderEventsBroker, usecase.NewAPIServerService, usecase.NewAppBuildService, usecase.NewAppDeployService, usecase.NewContinuousDeploymentService, usecase.NewRepositoryFetcherService, handlerSet,
	provideGRPCPort,
	provideWebServerConfig,
	provideImagePrefix,
	provideImageRegistry,
	provideRepositoryFetcherCacheDir, wire.FieldsOf(new(Config), "Builder", "SSGen", "DB", "MariaDB", "MongoDB"), wire.Struct(new(Router), "*"), wire.Bind(new(web.Router), new(*Router)), wire.Struct(new(Server), "*"),
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
