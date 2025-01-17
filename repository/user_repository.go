package repository

import (
	"go-rest-api-app/model"

	"gorm.io/gorm"
)

type IUserRepository interface {//インターフェースは、「こういう機能を持っているよ」という設計図です
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}


type userRepository struct {
	db *gorm.DB
}
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}
func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}