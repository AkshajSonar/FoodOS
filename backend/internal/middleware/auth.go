package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(
			tokenStr,
			claims,
			func(t *jwt.Token) (interface{}, error) {
				return []byte("super-secret-key"), nil
			},
		)
		if err != nil || claims["role"] != role {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}
