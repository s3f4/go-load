package app

import (
	"github.com/go-redis/redis/v8"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/repository"
	"gorm.io/gorm"
)

var mysqlConn *gorm.DB
var postgresConn *gorm.DB
var redisClient *redis.Client

func initConnections() {
	mysqlConn = repository.Connect(repository.MYSQL)
	postgresConn = repository.Connect(repository.POSTGRES)
	redisClient = repository.ConnectRedis()
}

func migrate() {
	mysqlConn.AutoMigrate(&models.Instance{})
	mysqlConn.AutoMigrate(&models.InstanceConfig{})
	mysqlConn.AutoMigrate(&models.TestGroup{})
	mysqlConn.AutoMigrate(&models.Test{})
	mysqlConn.AutoMigrate(&models.RunTest{})
	mysqlConn.AutoMigrate(&models.TransportConfig{})
	mysqlConn.AutoMigrate(&models.Header{})
	mysqlConn.AutoMigrate(&models.User{})
	mysqlConn.AutoMigrate(&models.Settings{})
}
