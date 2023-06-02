// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/grpc"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/repository"
	"github.com/traPtitech/neoshowcase/pkg/usecase"
)

// Injectors from wire.go:

func New(c2 Config) (*Server, error) {
	config := c2.DB
	db, err := repository.New(config)
	if err != nil {
		return nil, err
	}
	client, err := initBuildkitClient(c2)
	if err != nil {
		return nil, err
	}
	controllerServiceClientConfig := c2.Controller
	controllerBuilderServiceClient := grpc.NewControllerBuilderServiceClient(controllerServiceClientConfig)
	buildpackBackend, err := provideBuildpackBackend(c2)
	if err != nil {
		return nil, err
	}
	storageConfig := c2.Storage
	storage, err := provideStorage(storageConfig)
	if err != nil {
		return nil, err
	}
	publicKeys, err := provideRepositoryPublicKey(c2)
	if err != nil {
		return nil, err
	}
	imageConfig := c2.Image
	applicationRepository := repository.NewApplicationRepository(db)
	artifactRepository := repository.NewArtifactRepository(db)
	buildRepository := repository.NewBuildRepository(db)
	gitRepositoryRepository := repository.NewGitRepositoryRepository(db)
	builderService := usecase.NewBuilderService(controllerBuilderServiceClient, client, buildpackBackend, storage, publicKeys, imageConfig, applicationRepository, artifactRepository, buildRepository, gitRepositoryRepository)
	server := &Server{
		db:       db,
		buildkit: client,
		builder:  builderService,
	}
	return server, nil
}
