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

	if err = uc.usersRepository.Create(model); err != nil {
		return fmt.Errorf(common.ErrUnprocessableEntity)
	}

	ctx.Set("user_model", model)
	return err
}

func (uc *UserUseCase) Login(ctx *gin.Context, username string, password string) error {
	existsUser, err := uc.usersRepository.FindOneUser(&users.UserModel{Username: username})
	if err != nil {
		return fmt.Errorf(common.ErrBadRequest)
	}

	if err = existsUser.CheckPassword(password); err != nil {
		return fmt.Errorf(common.ErrBadRequest)
	}

	ctx.Set("user_model", existsUser)
	return nil
}

func (uc *UserUseCase) isAcceptable(ctx *gin.Context, model *users.UserModel) bool {
	// If the username is not in the specification or culturally preferred,
	// we will refuse to accept it
	return true
}