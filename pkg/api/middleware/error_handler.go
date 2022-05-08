package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pathak107/cloudesk/pkg/api/apierrors"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) > 0 {
			err := ctx.Errors[0].Err
			if apiErr, ok := err.(*apierrors.ApiError); ok {
				log.Println(apiErr.ErrorOriginal())
				ctx.JSON(apiErr.Code, gin.H{
					"error": apiErr.ErrMsg,
				})
				return
			} else {
				log.Println(err)
				ctx.JSON(-1, gin.H{
					"error": err.Error(),
				})
				return
			}
		}
	}
}
