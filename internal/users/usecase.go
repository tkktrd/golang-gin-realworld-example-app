package users

import "github.com/gin-gonic/gin"

type UseCase interface {
	Registration(ctx *gin.Context, model *UserModel) error
}