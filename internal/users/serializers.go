package users

import "github.com/gin-gonic/gin"

type UserResponse struct {
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Token    string  `json:"token"`
}

type UserSerializer interface {
	Response(ctx *gin.Context) (UserResponse, error)
}