package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"server/internal/config"
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
