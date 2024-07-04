package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/mineamihai2001/game-night/internal/infrastructure/opentelemetry"
)

func RequestIdMiddleware(ctx *gin.Context) {
	traceId := opentelemetry.GetTraceId(ctx.Request.Context())

	ctx.Header("x-request-id", traceId)
}
