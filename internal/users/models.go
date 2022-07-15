package users

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserModel struct {
	ID           uint    `gorm:"primary_key"`
	Username     string  `gorm:"column:username"`
	Email        string  `gorm:"column:email;unique_index"`
	PasswordHash string  `gorm:"column:hashed_password;not null"`
	CreatedAt time.Time  `gorm:"column:created_at;not null"`
	UpdatedAt time.Time  `gorm:"column:updated_at;not null"`
}

func (u UserModel) TableName() string {
	return "users"
}

func (u *UserModel) SetPassword(password string) error {
	if len(password) == 0 {
		return fmt.Errorf("password should not be empty")
	}
	bytePassword := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("password hashing failed, %v", err)
	}
	u.PasswordHash = string(passwordHash)
	return nil
}

func (u *UserModel) CheckPassword(password string) error {
	bytePassword := []byte(password)
	bytePasswordHash := []byte(u.PasswordHash)
	return bcrypt.CompareHashAndPassword(bytePasswordHash, bytePassword)
}