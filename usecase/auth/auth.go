package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	AccessTokenExpireDuration  = time.Minute * 5
	AccessTokenIssuer          = "mymyteam"
	ErrInvalidAccessToken      = "invalid access token"
	RefreshTokenExpireDuration = time.Hour * 2
	RefreshTokenIssuer         = "labulati"
	ErrInvalidRefreshToken     = "invalid refresh token"
)

var (
	AccessTokenSecert  = []byte("are you ok")
	RefreshTokenSecert = []byte("i am ok")
)

// Access token

type AccessTokenClaims struct {
	UserName string
	jwt.RegisteredClaims
}

func GenerateAccessToken(name string) (string, error) {
	claims := AccessTokenClaims{
		UserName: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessTokenExpireDuration)),
			Issuer:    AccessTokenIssuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(AccessTokenSecert)
}

func ParseAccessToken(token string) (*AccessTokenClaims, error) {
	jToken, err := jwt.ParseWithClaims(token, &AccessTokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return AccessTokenSecert, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := jToken.Claims.(*AccessTokenClaims); ok && jToken.Valid {
		return claims, nil
	}
	return nil, errors.New(ErrInvalidAccessToken)
}

// refresh token

type RefreshTokenClaims struct {
	UserName string
	jwt.RegisteredClaims
}

func GenerateRefreshToken(name string) (string, error) {
	claims := RefreshTokenClaims{
		UserName: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(RefreshTokenExpireDuration)),
			Issuer:    RefreshTokenIssuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(RefreshTokenSecert)
}

func ParseRefreshToken(token string) (*RefreshTokenClaims, error) {
	jToken, err := jwt.ParseWithClaims(token, &RefreshTokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return RefreshTokenSecert, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := jToken.Claims.(*RefreshTokenClaims); ok && jToken.Valid {
		return claims, nil
	}
	return nil, errors.New(ErrInvalidRefreshToken)
}
