package app

import (
	"github.com/s3f4/go-load/apigateway/handlers"
	"github.com/s3f4/go-load/apigateway/repository"
	"github.com/s3f4/go-load/apigateway/services"
)

var authHandler handlers.AuthHandler
var instanceHandler handlers.InstanceHandler
var runTestHandler handlers.RunTestHandler
var serviceHandler handlers.ServiceHandler
var statsHandler handlers.StatsHandler
var testGroupHandler handlers.TestGroupHandler
var testHandler handlers.TestHandler
var workerHandler handlers.WorkerHandler

func initHandlers() {
	authHandler = handlers.NewAuthHandler(
		repository.NewUserRepository(),
		services.NewAuthService(),
		services.NewTokenService(),
	)
	instanceHandler = handlers.NewInstanceHandler(services.NewInstanceService())
	runTestHandler = handlers.NewRunTestHandler(repository.NewRunTestRepository())
	serviceHandler = handlers.NewServiceHandler()
	statsHandler = handlers.NewStatsHandler(repository.NewResponseRepository())
	testGroupHandler = handlers.NewTestGroupHandler(repository.NewTestGroupRepository())
	testHandler = handlers.NewTestHandler(services.NewTestService(),
		repository.NewTestRepository())
	workerHandler = handlers.NewWorkerHandler()
}
