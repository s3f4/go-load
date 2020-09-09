package repository

import (
	"fmt"

	"github.com/s3f4/go-load/apigateway/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// BaseRepository an interface that uses sql
type BaseRepository interface {
	Insert(model interface{}) error
	Update(model interface{}) error
	Delete(model interface{}) error
	Get(model interface{}) error
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
func (br *baseRepository) connect() {
	dsn := "goload:go-load12345@tcp(mysql:3306)/go-load?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	br.DB = db
	// validations.RegisterCallbacks(br.DB)
	// br.DB.SetLogger(library.GetLogger())
	// br.DB.LogMode(config.GetBool("database.log_mode"))

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	//defer db.Close()
}

//Migrate migrates db
func (br *baseRepository) Migrate() {
	br.GetDB().AutoMigrate(&models.Instance{})
}

//GetDB return *gorm.DB instance
func (br *baseRepository) GetDB() *gorm.DB {
	if br.DB == nil {
		br.connect()
	}
	return br.DB
}

// Insert method
func (br *baseRepository) Insert(model interface{}) error {
	return br.GetDB().Create(model).Error
}

func (br *baseRepository) Update(model interface{}) error { return nil }
func (br *baseRepository) Delete(model interface{}) error { return nil }
func (br *baseRepository) Get(model interface{}) error    { return nil }
func (br *baseRepository) GetAll() error                  { return nil }
