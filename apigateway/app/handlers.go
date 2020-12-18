package app

import (
	"github.com/s3f4/go-load/apigateway/handlers"
	"github.com/s3f4/go-load/apigateway/repository"
	"github.com/s3f4/go-load/apigateway/services"
)

var authHandler handlers.AuthHandlerInterface

func initHandlers() {
	authHandler = handlers.NewAuthHandler(
		repository.NewUserRepository(),
		services.NewAuthService(),
		services.NewTokenService(),
	)
}
