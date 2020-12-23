package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-redis/redis/v8"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DBType database type for connection different types of databases
type DBType string

const (
	// MYSQL DBType
	MYSQL DBType = "MYSQL"
	// POSTGRES DBType
	POSTGRES = "POSTGRES"
)

const (
	// PostgresDSN postgres connection string
	PostgresDSN = "user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=UTC"
	// MySQLDSN mysql connection string
	MySQLDSN = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
)

// Connect db
func Connect(dbType DBType) *gorm.DB {
	var dsnStr string

	if dbType == MYSQL {
		dsnStr = MySQLDSN
	} else if dbType == POSTGRES {
		dsnStr = PostgresDSN
	}

	dsn := fmt.Sprintf(dsnStr,
		os.Getenv(dsnStr+"_USER"),
		os.Getenv(dsnStr+"_PASSWORD"),
		os.Getenv(dsnStr+"_HOST"),
		os.Getenv(dsnStr+"_PORT"),
		os.Getenv(dsnStr+"_DATABASE"),
	)

	var dialector gorm.Dialector

	if dbType == MYSQL {
		dialector = mysql.Open(dsn)
	} else if dbType == POSTGRES {
		dialector = postgres.Open(dsn)
	}

	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: getLogger(),
	})

	if err != nil {
		log.Panicf("failed to connect database: %s", err)
	}

	return db
}

// ConnectMock mock
func ConnectMock() (*sql.DB, sqlmock.Sqlmock, *gorm.DB) {
	sqlDB, sqlMock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      sqlDB,
	}), &gorm.Config{})

	if err != nil {
		log.Panicf("failed to connect mock database: %s", err)
	}

	return sqlDB, sqlMock, db
}

func getLogger() logger.Interface {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)
	return newLogger
}

// ConnectRedis connect redis
func ConnectRedis() *redis.Client {
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "redis:6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr: dsn,
		DB:   0,
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		panic(err)
	}
	return client
}
