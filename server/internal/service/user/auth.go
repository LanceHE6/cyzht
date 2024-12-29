package user

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	jwt2 "server/pkg/jwt"
	"server/pkg/response"
	"strings"
)

// AuthMiddleware 基础鉴权的中间件
func (s userService) AuthMiddleware() gin.HandlerFunc {
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
		myClaims, err := GetUserInfoByContext(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, response.NewResponse(13, "非法token", err.Error()))
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(bearerToken[1], &jwt2.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
			return jwt2.JwtKey, nil
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

		if _, ok := token.Claims.(*jwt2.MyClaims); ok && token.Valid {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, response.NewResponse(18, "非法token", nil))
			c.Abort()
			return
		}
	}
}

// GetUserInfoByContext
//
//	@Description: 从context中获取用户信息
//	@param context *gin.Context
//	@return pkg.MyClaims 用户信息
//	@return error 错误信息
func GetUserInfoByContext(context *gin.Context) (jwt2.MyClaims, error) {
	authHeader := context.GetHeader("Authorization")
	bearerToken := strings.Split(authHeader, " ")
	// 解析token
	claims := jwt2.MyClaims{}
	_, err := jwt.ParseWithClaims(bearerToken[1], &claims, func(token *jwt.Token) (interface{}, error) {
		return jwt2.JwtKey, nil
	})
	// 从token中获取载荷数据
	return claims, err
}

// Auth
//
//	@Description: token验证
//	@param token string
//	@return bool 是否验证成功
//	@return pkg.MyClaims 用户信息
func (s userService) Auth(token string) (bool, jwt2.MyClaims) {
	// 解析token
	claims := jwt2.MyClaims{}
	bearToken, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return jwt2.JwtKey, nil
	})

	if err != nil {
		return false, claims
	}

	// 判断是否在数据库中
	user := s.UserRepo.SelectByID(claims.ID)
	if user == nil {
		return false, claims
	}
	if user.SessionID != claims.SessionID {
		return false, claims
	}

	if _, ok := bearToken.Claims.(*jwt2.MyClaims); ok && bearToken.Valid {
		return true, claims
	} else {
		return false, claims
	}
}
