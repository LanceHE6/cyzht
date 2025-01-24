package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"server/internal/config"
	"server/internal/repo"
	"strings"
	"time"
)

var SecretKey = []byte(config.GetConfig().Server.JWTSecret) // 用于签名的密钥

func KeyFunc(*jwt.Token) (interface{}, error) {
	return SecretKey, nil
}

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
	tokenString, err := token.SignedString(SecretKey)

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
	claims, ok := Check(bearerToken[1])
	if !ok {
		return nil, errors.New("invalid token")
	}
	return &claims, nil
}

// Check
//
//	@Description: token验证
//	@param token string
//	@return bool 是否通过
//	@return MyClaims 载荷数据
func Check(tokenStr string) (MyClaims, bool) {
	// 解析token
	claims := MyClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, &claims, KeyFunc)

	if err != nil {
		return claims, false
	}

	// 校验token的session_id是否有效
	userRepo := repo.GetRepo().UserRepo
	user := userRepo.SelectByID(claims.ID)
	if user == nil || user.SessionID != claims.SessionID {
		return claims, false
	}

	if _, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, true
	} else {
		return claims, false
	}
}
