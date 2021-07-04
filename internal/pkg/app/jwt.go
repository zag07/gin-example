package app

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/cast"
	"github.com/zs368/gin-example/configs"
	"github.com/zs368/gin-example/internal/pkg/utils"
)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GetJWTSecret() []byte {
	return []byte(configs.Auth.JwtSecret)
}

func GenerateToken(Username, Password string) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Username: utils.EncodeMD5(Username),
		Password: utils.EncodeMD5(Password),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(cast.ToDuration(configs.Auth.JwtExpire)).Unix(),
			Issuer:    configs.Auth.JwtIssuer,
		},
	}).SignedString(GetJWTSecret())

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})

	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
