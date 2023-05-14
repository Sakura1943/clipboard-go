// Token 生成与解析
package token

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// JWT密钥
var JwtSecret = []byte("sakunia")

// Token携带参数结构体
type Claims struct {
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
	Permission   string `json:"permission"`
	jwt.StandardClaims
}

// Token生成
func GenerateToken(UserName string, UserPassword string, Permission string) (string, error) {
	// 获取当前时间
	nowTime := time.Now()
	// 有效期(当前时间往后的3小时)
	expireTime := nowTime.Add(3 * time.Hour)

	// 定义Token携带的信息
	claims := Claims{
		UserName,
		UserPassword,
		Permission,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	// 获取Token
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(JwtSecret)
	return token, err
}

// 解析Token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})

	// token解析不为空则获取成功
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	// 否则获取失败
	return nil, err
}
