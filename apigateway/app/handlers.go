package app

import (
	"github.com/s3f4/go-load/apigateway/handlers"
	"github.com/s3f4/go-load/apigateway/middlewares"
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
var m *middlewares.Middleware

func initHandlers() {
	userRepository := repository.NewUserRepository()
	runTestRepository := repository.NewRunTestRepository()
	testRepository := repository.NewTestRepository()
	testGroupRepository := repository.NewTestGroupRepository()
	responseRepository := repository.NewResponseRepository()
	redisRepository := repository.NewRedisRepository()
	instanceRepository := repository.NewInstanceRepository()

	queue := services.NewRabbitMQService()
	authService := services.NewAuthService(redisRepository)
	tokenService := services.NewTokenService()
	instanceService := services.NewInstanceService(instanceRepository)
	testService := services.NewTestService(
		instanceRepository,
		testRepository,
		runTestRepository,
		queue,
	)

	m = middlewares.NewMiddleware(
		tokenService,
		authService,
		testRepository,
		testGroupRepository,
		runTestRepository,
	)

	authHandler = handlers.NewAuthHandler(
		userRepository,
		authService,
		tokenService,
	)
	instanceHandler = handlers.NewInstanceHandler(instanceService)
	runTestHandler = handlers.NewRunTestHandler(runTestRepository)
	serviceHandler = handlers.NewServiceHandler()
	statsHandler = handlers.NewStatsHandler(responseRepository)
	testGroupHandler = handlers.NewTestGroupHandler(testGroupRepository)
	testHandler = handlers.NewTestHandler(testService, testRepository)
	workerHandler = handlers.NewWorkerHandler()
}
