package common

import (
	"errors"
	"personal_page/zero/user/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	accessExpire         = time.Minute * 5
	refreshExpire        = time.Hour * 24
	accessSecret  []byte = []byte("hello")
	refreshSecret []byte = []byte("thank you")
)

type UserClaims struct {
	Username string
	jwt.RegisteredClaims
}

func Config(conf *config.Config) {
	accessExpire = time.Duration(conf.Auth.AccessExpire) * time.Second
	accessSecret = []byte(conf.Auth.AccessSecret)
	refreshExpire = time.Duration(conf.Auth.RefreshExpire) * time.Second
	refreshSecret = []byte(conf.Auth.RefreshSecret)
}

func GenerateAccessToken(username string) (string, error) {
	claims := UserClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessExpire)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(accessSecret)
}

// // 中间件会自动解析token  这里不需要提供解析方法
// func ParseAccessToken(signedStr string) (*UserClaims, error) {
// 	token, err := jwt.ParseWithClaims(signedStr, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
// 		return AccessTokenSecret, nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	if claims, ok := token.Claims.(*UserClaims); ok {
// 		return claims, nil
// 	}
// 	return nil, errors.New("invalid token")
// }

func GenerateRefreshToken(username string) (string, error) {
	claims := UserClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshExpire)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(refreshSecret)
}

func ParseRefreshToken(signedStr string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(signedStr, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return refreshSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*UserClaims); ok {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
