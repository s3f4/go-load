package repository

import (
	"github.com/s3f4/go-load/apigateway/models"
	"gorm.io/gorm"
)

// UserRepository for auth db
type UserRepository interface {
	DB() *gorm.DB
	Register(*models.User) error
	Login(*models.User) error
	Logout(*models.User) error
	Get(id uint) (*models.User, error)
}

type userRepository struct {
	base BaseRepository
}

var userRepositoryObject UserRepository

// NewUserRepository returns an testRepository object
func NewUserRepository() UserRepository {
	if userRepositoryObject == nil {
		userRepositoryObject = &userRepository{
			base: NewBaseRepository(MYSQL),
		}
	}
	return userRepositoryObject
}

func (r *userRepository) DB() *gorm.DB {
	return r.base.GetDB()
}

func (r *userRepository) Register(user *models.User) error {
	return nil
}

func (r *userRepository) Login(user *models.User) error {
	return nil
}

func (r *userRepository) Logout(user *models.User) error {
	return nil
}

func (r *userRepository) Get(id uint) (*models.User, error) {
	return nil, nil
}
