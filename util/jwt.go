package util

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

var MySecret = []byte("Let's Go Home")

const (
	TokenMaxExpireHour = 1 // token最长有效期
)

type MyClaims struct {
	UserId   int64
	UserName string
	Password string
	jwt.RegisteredClaims
}

func CreateToken(userid int64, username string, password string) (string, error) {
	// log.Println("调用了 CreateToken")
	claim := MyClaims{
		UserId:   userid,
		UserName: username,
		Password: password,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now()),                                     // 生效
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenMaxExpireHour * time.Hour)), // 失效
			Issuer:    "GoHome",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	s, err := token.SignedString(MySecret)
	if err != nil {
		log.Println("Token 生成错误:", err)
		return "", err
	}

	// log.Println("生成的 Token:", s)

	return s, nil
}

func ParseToken(s string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(s, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		log.Println("Token 解析错误:", err)
		return nil, err
	}
	// log.Println("解析后的 Token:", token)
	if !token.Valid {
		return nil, errors.New("claim invalid")
	}
	claims, ok := token.Claims.(*MyClaims)
	if !ok {
		return nil, errors.New("invalid claim type")
	}
	return claims, nil
}
