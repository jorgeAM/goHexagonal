package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	}
}
