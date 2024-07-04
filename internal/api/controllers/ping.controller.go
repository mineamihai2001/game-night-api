package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type PingController struct {
	context       context.Context
	cancelContext context.CancelFunc
}

func NewPingController() *PingController {
	c := &PingController{}
	c.context, c.cancelContext = c.createContext()

	return c
}

func (c *PingController) createContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

func (c *PingController) Ping(ctx *gin.Context) {
	log.
		Info().
		Ctx(ctx.Request.Context()).
		Msg("Request received")

	ctx.JSON(http.StatusOK, gin.H{
		"message":     "pong",
		"status":      "up",
		"retrievedAt": time.Now(),
	})
}
