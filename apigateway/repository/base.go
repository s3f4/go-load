package repository

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/qor/validations"
	"github.com/s3f4/go-load/apigateway/models"
)

// BaseRepository an interface that uses sql
type BaseRepository interface {
	Init()
}

type baseRepository struct {
	db *gorm.DB
}

// NewBaseRepository instance of baseRepository
func NewBaseRepository() BaseRepository {
	return &baseRepository{db: nil}
}

//InitDatabase initializes a global database instance
func (br *baseRepository) Init() {
	connectionString := fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"password",
		"go_load")

	var err error
	db, err := gorm.Open("mysql", connectionString)
	validations.RegisterCallbacks(db)
	// db.SetLogger(library.GetLogger())
	// db.LogMode(config.GetBool("database.log_mode"))

	db.Set("gorm:association_autoupdate", false)
	db.Set("gorm:association_autocreate", false)

	if err != nil {
		panic("failed to connect database")
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.DB().SetMaxIdleConns(20)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	db.DB().SetConnMaxLifetime(time.Minute * 5)

	//defer db.Close()
	br.db = db
}

func (br *baseRepository) DropTableIfExists() {
	br.db.DropTableIfExists(&models.Instance{})
}

//Migrate migrates db
func (br *baseRepository) Migrate() {
	br.db.AutoMigrate(&models.Instance{})
}

//GetDB return *gorm.DB instance
func (br *baseRepository) GetDB() *gorm.DB {
	if br.db == nil || br == nil {
		br.Init()
	}
	return br.db
}
