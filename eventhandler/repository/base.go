package repository

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)



// BaseRepository an interface that uses sql
type BaseRepository interface {
	GetDB() *gorm.DB
	Migrate(models ...interface{}) error
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

//GetDB return *gorm.DB instance
func (r *baseRepository) GetDB() *gorm.DB {
	if r.DB == nil {
		r.connect()
	}
	return r.DB
}

//Migrate migrates db
func (r *baseRepository) Migrate(models ...interface{}) error {
	for _, model := range models {
		err := r.GetDB().AutoMigrate(model)
		if err != nil {
			return err
		}
	}
	return nil
}
