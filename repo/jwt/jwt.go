package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 自定义Claims结构（携带的数据）
type MyCustomClaims struct {
	UserID               string `json:"user_id"`
	jwt.RegisteredClaims        // 标准Claims（过期时间等）
}

// 密钥（务必保密！）
var secretKey = []byte("your_secret_key_here")

// 生成JWT
func generateJWT(userID string) (string, error) {
	claims := MyCustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 签发时间
			Issuer:    "my_app",                                           // 签发者
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// 验证并解析JWT
func verifyJWT(tokenString string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func main() {
	// 生成Token
	token, err := generateJWT("user123")
	if err != nil {
		panic(err)
	}
	fmt.Println("Generated Token:\n", token)

	// 验证Token
	claims, err := verifyJWT(token)
	if err != nil {
		fmt.Println("Verification failed:", err)
		return
	}
	fmt.Printf("Valid Token! UserID: %s, ExpiresAt: %v\n",
		claims.UserID, claims.ExpiresAt.Time.Format(time.RFC3339))
}
