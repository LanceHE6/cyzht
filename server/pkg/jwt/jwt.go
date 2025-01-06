package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"server/internal/config"
	"strings"
	"time"
)

var JwtKey = []byte(config.ConfigData.Server.JWTSecret) // 用于签名的密钥

// MyClaims 自定义载荷内容
type MyClaims struct {
	jwt.StandardClaims
	UserClaims
}

type UserClaims struct {
	ID        int64  `json:"id"`
	SessionID string `json:"session_id"`
}

// GenerateToken
//
//	@Description: 生成一个token
//	@param id 用户id
//	@param permission 用户权限
//	@param createdAt 用户创建时间
//	@param sessionID 用户sessionID
//	@return string token
//	@return error 错误
func GenerateToken(id int64, sessionID string) (string, error) {
	expirationTime := time.Now().Add(72 * time.Hour)
	claims := &MyClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
		UserClaims: UserClaims{
			ID:        id,
			SessionID: sessionID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// GetClaimsByContext
//
//	@Description: 从context中获取用户信息
//	@param context *gin.Context
//	@return pkg.MyClaims 用户信息
//	@return error 错误信息
func GetClaimsByContext(context *gin.Context) (*MyClaims, error) {
	authHeader := context.GetHeader("Authorization")
	bearerToken := strings.Split(authHeader, " ")
	// 解析token
	claims := &MyClaims{}
	_, err := jwt.ParseWithClaims(bearerToken[1], claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	// 从token中获取载荷数据
	return claims, err
}

// Check
//
//	@Description: token验证
//	@param token string
//	@return bool 是否通过
//	@return MyClaims 载荷数据
func Check(token string) (MyClaims, bool) {
	// 解析token
	claims := MyClaims{}
	bearToken, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil {
		return claims, false
	}

	if _, ok := bearToken.Claims.(*MyClaims); ok && bearToken.Valid {
		return claims, true
	} else {
		return claims, false
	}
}
