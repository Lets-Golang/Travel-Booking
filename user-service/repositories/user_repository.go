package repositories

import (
	"github.com/Lets-Golang/Travel-Booking/user-service/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.UserEntity) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByID(id int) (*models.UserEntity, error) {
	var user models.UserEntity
	err := r.db.First(&user, id).Error
	return &user, err
}
