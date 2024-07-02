package middleware

import (
	"github.com/gin-gonic/gin"
)

func Query[T interface{}](ctx *gin.Context) (*T, error) {
	var query T

	if err := ctx.ShouldBindQuery(&query); err != nil {
		return nil, err
	}

	return &query, nil
}
