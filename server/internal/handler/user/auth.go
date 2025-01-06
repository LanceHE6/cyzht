package user

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	myjwt "server/pkg/jwt"
	"server/pkg/response"
	"strings"
)

// AuthMiddleware 基础鉴权的中间件
func (s userHandler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, response.NewResponse(11, "未提供 Authorization 鉴权头", nil))
			c.Abort()
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			c.JSON(http.StatusUnauthorized, response.NewResponse(12, "非法token", nil))
			c.Abort()
			return
		}
		myClaims, err := myjwt.GetClaimsByContext(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, response.NewResponse(13, "非法token", err.Error()))
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
					c.JSON(http.StatusUnauthorized, response.NewResponse(14, "token已过期", nil))
				} else {
					// Other errors
					c.JSON(http.StatusUnauthorized, response.NewResponse(15, "非法token", nil))
				}
			}
			c.Abort()
			return
		}

		// 判断是否在数据库中

		user := s.UserRepo.SelectByID(myClaims.ID)
		if user == nil {
			c.JSON(http.StatusUnauthorized, response.NewResponse(16, "用户不存在", nil))
			c.Abort()
		}
		if user.SessionID != myClaims.SessionID {
			c.JSON(http.StatusUnauthorized, response.NewResponse(17, "token已失效", nil))
			c.Abort()
		}

		if _, ok := token.Claims.(*myjwt.MyClaims); ok && token.Valid {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, response.NewResponse(18, "非法token", nil))
			c.Abort()
			return
		}
	}
}
