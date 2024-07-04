package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mineamihai2001/game-night/internal/api/controllers"
	"github.com/mineamihai2001/game-night/internal/api/middleware"
	"github.com/mineamihai2001/game-night/pkg/tracing"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func Create() *gin.Engine {
	app := gin.Default()

	// auto instrument requests
	app.Use(otelgin.Middleware("game-night-api", otelgin.WithFilter(tracing.IgnorePaths)))
	app.Use(middleware.RequestIdMiddleware)

	v1 := app.Group("/v1")

	root := v1.Group("/")
	{
		pingController := controllers.NewPingController()
		root.GET("/ping", pingController.Ping)
	}

	return app
}
