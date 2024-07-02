package middleware

import "github.com/gin-gonic/gin"

func Body[T interface{}](ctx *gin.Context) (*T, error) {
	var body T

	if err := ctx.ShouldBind(&body); err != nil {
		return nil, err
	}

	return &body, nil
}
