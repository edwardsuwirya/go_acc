package middleware

import (
	"enigmacamp.com/goacc/delivery/commonresp"
	"enigmacamp.com/goacc/logger"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detectedError := c.Errors.Last()
		if detectedError == nil {
			return
		}
		e := detectedError.Error()
		logger.Log.Error().Err(fmt.Errorf("%v", e)).Msg("User Error")
		commonresp.NewErrorJsonResponse(c, detectedError).Send()
	}
}
