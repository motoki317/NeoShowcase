// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/leandro-lugaresi/hub"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/admindb"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/eventbus"
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/web"
	"github.com/traPtitech/neoshowcase/pkg/interface/handler"
	"github.com/traPtitech/neoshowcase/pkg/interface/repository"
	"github.com/traPtitech/neoshowcase/pkg/usecase"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func New(c2 Config) (*Server, error) {
	hubHub := hub.New()
	bus := eventbus.NewLocal(hubHub)
	config := provideAdminDBConfig()
	db, err := admindb.New(config)
	if err != nil {
		return nil, err
	}
	webhookSecretRepository := repository.NewWebhookSecretRepository(db)
	gitPushWebhookService := usecase.NewGitPushWebhookService(webhookSecretRepository)
	webhookReceiverHandler := handler.NewWebhookReceiverHandler(bus, gitPushWebhookService)
	router := &Router{
		wr: webhookReceiverHandler,
	}
	webConfig := provideWebServerConfig(router)
	server := web.NewServer(webConfig)
	mainServer, err := NewServer(c2, server, db, hubHub)
	if err != nil {
		return nil, err
	}
	return mainServer, nil
}
