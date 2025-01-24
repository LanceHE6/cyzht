package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/pkg/jwt"
	"server/pkg/response"
	"strings"
)

// Auth 基础鉴权中间件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, response.NewResponse(11, "Authorization header not provided", nil))
			c.Abort()
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || (bearerToken[0] != "Bearer" && bearerToken[0] != "bearer") {
			c.JSON(http.StatusUnauthorized, response.NewResponse(12, "invalid Authorization header format", nil))
			c.Abort()
			return
		}

		if _, ok := jwt.Check(bearerToken[1]); !ok {
			c.JSON(http.StatusUnauthorized, response.NewResponse(13, "invalid token", nil))
			c.Abort()
			return
		}
		c.Next()
	}
}
