package repository

import (
	"github.com/s3f4/go-load/apigateway/models"
	"gorm.io/gorm"
)

// UserRepository for auth db
type UserRepository interface {
	DB() *gorm.DB
	Create(*models.User) error
	GetByEmailAndPassword(*models.User) (*models.User, error)
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

func (r *userRepository) Create(user *models.User) error {
	return r.DB().Create(user).Error
}

func (r *userRepository) GetByEmailAndPassword(user *models.User) (*models.User, error) {
	var dbUser models.User
	if err := r.DB().Where("email=? AND password=?", user.Email, user.Password).Take(&dbUser).Error; err != nil {
		return nil, err
	}
	return &dbUser, nil
}
