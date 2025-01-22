package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	myjwt "server/pkg/jwt"
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

		token, err := jwt.ParseWithClaims(bearerToken[1], &myjwt.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
			return myjwt.JwtKey, nil
		})

		if err != nil {
			var ve *jwt.ValidationError
			if errors.As(err, &ve) {
				if ve.Errors&jwt.ValidationErrorExpired != 0 {
					// Token is expired
					c.JSON(http.StatusUnauthorized, response.NewResponse(13, "token is expired", nil))
				} else {
					// Other errors
					c.JSON(http.StatusUnauthorized, response.NewResponse(14, "invalid token", nil))
				}
			}
			c.Abort()
			return
		}
		if _, ok := token.Claims.(*myjwt.MyClaims); ok && token.Valid {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, response.NewResponse(15, "invalid token", nil))
			c.Abort()
			return
		}
	}
}
