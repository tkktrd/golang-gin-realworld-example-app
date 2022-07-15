package users

import "github.com/gin-gonic/gin"

type UseCase interface {
	Login(ctx *gin.Context, username string, password string) error
	Registration(ctx *gin.Context, model *UserModel) error
}