package usecase

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tkktrd/golang-gin-realworld-example-app/internal/users"
	"tkktrd/golang-gin-realworld-example-app/pkg/common"
)

type UserUseCase struct {
	usersRepository users.Repository
}

func NewUserUseCase(usersRepository users.Repository) users.UseCase {
	return &UserUseCase{usersRepository: usersRepository}
}

func (uc *UserUseCase) Registration(ctx *gin.Context, model *users.UserModel) error {
	existsUser, err := uc.usersRepository.FindOneUser(&users.UserModel{Username: model.Username})
	if existsUser != nil || err == nil {
		return fmt.Errorf(common.ErrUserIdAlreadyExists)
	}

	if !uc.isAcceptable(ctx, model) {
		return fmt.Errorf(common.ErrBadRequest)
	}

	err = uc.usersRepository.Create(model)
	return err
}

func (uc *UserUseCase) isAcceptable(ctx *gin.Context, model *users.UserModel) bool {
	// If the username is not in the specification or culturally preferred,
	// we will refuse to accept it
	return true
}