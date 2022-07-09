package usecase

import (
	"github.com/gin-gonic/gin"
	"tkktrd/golang-gin-realworld-example-app/internal/users"
)

type UserUseCase struct {
	usersRepository users.Repository
}

func NewUserUseCase(usersRepository users.Repository) users.UseCase {
	return &UserUseCase{usersRepository: usersRepository}
}

func (uc *UserUseCase) Registration(ctx *gin.Context, model *users.UserModel) error {
	err := uc.usersRepository.Create(model)
	return err
}