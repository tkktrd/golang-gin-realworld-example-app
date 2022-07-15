package users

import "github.com/gin-gonic/gin"

type Handlers interface {
	Registration() gin.HandlerFunc
	Login() gin.HandlerFunc
}