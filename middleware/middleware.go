package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project/tokens"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//token := c.GetHeader("Authorization")
		//if token == "" {
		//	c.AbortWithStatus(http.StatusUnauthorized)
		//}
		err := tokens.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
