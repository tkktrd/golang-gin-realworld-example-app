package v1

import (
	"github.com/gin-gonic/gin"
	"tkktrd/golang-gin-realworld-example-app/internal/users"
)

func MapNewRoutes(v1 *gin.RouterGroup, h users.Handlers) {
	v1.POST("/users", h.Registration())
	v1.POST("/user/login", h.Login())
}