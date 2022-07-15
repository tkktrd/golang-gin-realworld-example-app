package repository

import (
	"gorm.io/gorm"
	"tkktrd/golang-gin-realworld-example-app/internal/users"
)

type UsersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) users.Repository {
	return &UsersRepository{db: db}
}

func (repo *UsersRepository) Create(user *users.UserModel) error {
	err := repo.db.Save(user).Error
	return err
}

func (repo *UsersRepository) FindOneUser(condition interface{}) (*users.UserModel, error) {
	var model users.UserModel
	if err := repo.db.Where(condition).First(&model).Error; err != nil {
		return nil, err
	}
	return &model, nil
}