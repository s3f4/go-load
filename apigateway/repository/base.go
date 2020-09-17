package repository

import (
	"fmt"
	"os"

	"github.com/s3f4/mu/log"

	"github.com/s3f4/go-load/apigateway/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// BaseRepository an interface that uses sql
type BaseRepository interface {
	Insert(model interface{}) error
	Update(model interface{}) error
	Delete(model interface{}) error
	Get(model interface{}, query ...interface{}) error
	GetAll() error
	Migrate()
}

type baseRepository struct {
	*gorm.DB
}

// NewBaseRepository instance of baseRepository
func NewBaseRepository() BaseRepository {
	return &baseRepository{}
}

//InitDatabase initializes a global database instance
func (r *baseRepository) connect() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	r.DB = db
	// validations.RegisterCallbacks(br.DB)
	// br.DB.SetLogger(library.GetLogger())
	// br.DB.LogMode(config.GetBool("database.log_mode"))

	if err != nil {
		log.Panicf("failed to connect database: %s", err)
	}

	//defer db.Close()
}

//Migrate migrates db
func (r *baseRepository) Migrate() {
	r.GetDB().AutoMigrate(&models.Instance{})
}

//GetDB return *gorm.DB instance
func (r *baseRepository) GetDB() *gorm.DB {
	if r.DB == nil {
		r.connect()
	}
	return r.DB
}

// Insert method
func (r *baseRepository) Insert(model interface{}) error {
	return r.GetDB().Create(model).Error
}

func (r *baseRepository) Update(model interface{}) error { return nil }
func (r *baseRepository) Delete(model interface{}) error { return nil }
func (r *baseRepository) Get(model interface{}, query ...interface{}) error {
	if query == nil {
		return r.GetDB().Last(model).Error
	}
	return nil
}
func (r *baseRepository) GetAll() error { return nil }
