package repository

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/s3f4/go-load/apigateway/models"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DBType database type for connection different types of databases
type DBType int8

const (
	// MYSQL DBType
	MYSQL DBType = iota
	// POSTGRES DBType
	POSTGRES
)

type connect struct {
	PostgresDSN string
	MySQLDSN    string
}

// newConnection config
func newConnection() *connect {
	return &connect{
		PostgresDSN: "user=%s password=%s host=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		MySQLDSN:    "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	}
}

func (c *connect) getLogger() logger.Interface {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // Disable color
		},
	)
	return newLogger
}

func (c *connect) connectMYSQL(r *baseRepository) {
	dsn := fmt.Sprintf(c.MySQLDSN,
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: c.getLogger(),
	})
	r.DB = db
	if err != nil {
		log.Panicf("failed to connect database: %s", err)
	}
}

func (c *connect) connectPOSTGRES(r *baseRepository) {
	dsn := fmt.Sprintf(c.PostgresDSN,
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_DATABASE"),
		os.Getenv("POSTGRES_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: c.getLogger(),
	})
	r.DB = db
	if err != nil {
		log.Panicf("failed to connect database: %s", err)
	}
}

// BaseRepository an interface that uses sql
type BaseRepository interface {
	Migrate()
	GetDB() *gorm.DB
}

type baseRepository struct {
	*gorm.DB
	dbType DBType
}

// NewBaseRepository instance of baseRepository
func NewBaseRepository(dbType DBType) BaseRepository {
	return &baseRepository{
		dbType: dbType,
	}
}

// connect initializes a global database instance
func (r *baseRepository) connect() {
	c := newConnection()
	switch r.dbType {
	case MYSQL:
		c.connectMYSQL(r)
	case POSTGRES:
		c.connectPOSTGRES(r)
	}

	// validations.RegisterCallbacks(br.DB)
	// br.DB.SetLogger(library.GetLogger())
	// br.DB.LogMode(config.GetBool("database.log_mode"))

	//defer db.Close()
}

//Migrate migrates db
func (r *baseRepository) Migrate() {
	r.GetDB().AutoMigrate(&models.Instance{})
	r.GetDB().AutoMigrate(&models.InstanceConfig{})
}

//GetDB return *gorm.DB instance
func (r *baseRepository) GetDB() *gorm.DB {
	if r.DB == nil {
		r.connect()
	}
	return r.DB
}
