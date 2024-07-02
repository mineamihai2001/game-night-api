package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mineamihai2001/game-night/internal/api/controllers"
)

func Create() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")

	gen := v1.Group("/")
	{
		pingController := controllers.NewPingController()
		gen.GET("/ping", pingController.Ping)
	}

	return router
}
