package v1

import (
	"github.com/gin-gonic/gin"
	"tkktrd/golang-gin-realworld-example-app/internal/users"
)

func MapNewRoutes(v1 *gin.RouterGroup, h users.Handlers) {
	r := v1.Group("/users")
	r.POST("/", h.Registration())
	r.POST("/login", h.Login())
}