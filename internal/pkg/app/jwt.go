package app

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// TODO 后面改掉 不能存敏感信息到 token 中
type UserInfo struct {
	Uid      uint   `json:"uid"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type Claims struct {
	*jwt.StandardClaims
	UserInfo
}

func GetJWTSecret() []byte {
	return []byte("")
}

func GenerateToken(user UserInfo) (string, error) {
	t := jwt.New(jwt.GetSigningMethod("HS256"))

	t.Claims = &Claims{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7200).Unix(),
			Issuer:    "",
		},
		user,
	}

	return t.SignedString(GetJWTSecret())
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
