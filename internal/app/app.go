package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Run() {
	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
			"message": "Hello World",
		})
	})
	engine.Run(":3000")
}