package users

import "github.com/gin-gonic/gin"

type Validators interface {
	GetUsername() string
	GetPassword() string
	GetUserModel() *UserModel
	Bind(c *gin.Context) error
}