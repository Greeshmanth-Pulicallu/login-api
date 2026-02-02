package router

import (
	"github.com/gin-gonic/gin"

	"github.com/Greeshmanth-Pulicallu/login-api/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/login", handlers.LoginHandler)

	return r
}
