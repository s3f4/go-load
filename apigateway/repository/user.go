package repository

import (
	"github.com/s3f4/go-load/apigateway/models"
	"gorm.io/gorm"
)

// UserRepository for auth db
type UserRepository interface {
	Create(*models.User) error
	Get(ID uint) (*models.User, error)
	GetByEmailAndPassword(*models.User) (*models.User, error)
	List() ([]*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository returns an testRepository object
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) Get(ID uint) (*models.User, error) {
	var dbUser models.User
	if err := r.db.First(&dbUser, ID).Error; err != nil {
		return nil, err
	}
	return &dbUser, nil
}

func (r *userRepository) GetByEmailAndPassword(user *models.User) (*models.User, error) {
	var dbUser models.User
	if err := r.db.Where("email=? AND password=?", user.Email, user.Password).Take(&dbUser).Error; err != nil {
		return nil, err
	}
	return &dbUser, nil
}

func (r *userRepository) List() ([]*models.User, error) {
	var users []*models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
