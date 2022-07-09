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
