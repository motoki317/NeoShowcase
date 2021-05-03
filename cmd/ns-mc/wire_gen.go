// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/traPtitech/neoshowcase/pkg/infrastructure/web"
	"github.com/traPtitech/neoshowcase/pkg/interface/handler"
	"github.com/traPtitech/neoshowcase/pkg/usecase"
)

// Injectors from wire.go:

func NewServer() (*web.Server, error) {
	trapShowcaseJWTPublicKeyPEM, err := providePubKeyPEM()
	if err != nil {
		return nil, err
	}
	memberCheckService, err := usecase.NewMemberCheckService(trapShowcaseJWTPublicKeyPEM)
	if err != nil {
		return nil, err
	}
	memberCheckHandler := handler.NewMemberCheckHandler(memberCheckService)
	router := &Router{
		h: memberCheckHandler,
	}
	config := provideServerConfig(router)
	server := web.NewServer(config)
	return server, nil
}
