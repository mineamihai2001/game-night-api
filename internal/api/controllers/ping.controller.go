package controllers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
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
	ctx.JSON(200, gin.H{
		"message":     "pong",
		"status":      "up",
		"retrievedAt": time.Now(),
	})
}
