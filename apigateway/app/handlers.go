package app

import (
	"github.com/go-redis/redis/v8"
	"github.com/s3f4/go-load/apigateway/handlers"
	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/middlewares"
	"github.com/s3f4/go-load/apigateway/repository"
	"github.com/s3f4/go-load/apigateway/services"
	"gorm.io/gorm"
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
var mysqlConn *gorm.DB
var postgresConn *gorm.DB
var redisClient *redis.Client

func initHandlers() {
	command := library.NewCommand()
	mysqlConn = repository.Connect(repository.MYSQL)
	postgresConn = repository.Connect(repository.POSTGRES)
	redisClient = repository.ConnectRedis()

	userRepository := repository.NewUserRepository(mysqlConn)
	runTestRepository := repository.NewRunTestRepository(mysqlConn)
	testRepository := repository.NewTestRepository(mysqlConn)
	testGroupRepository := repository.NewTestGroupRepository(mysqlConn)
	responseRepository := repository.NewResponseRepository(postgresConn)
	redisRepository := repository.NewRedisRepository(redisClient)
	instanceRepository := repository.NewInstanceRepository(mysqlConn, command)

	queue := services.NewQueueService()
	authService := services.NewAuthService(redisRepository)
	tokenService := services.NewTokenService()
	instanceService := services.NewInstanceService(instanceRepository, command)
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
