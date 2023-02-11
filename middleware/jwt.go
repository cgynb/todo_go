package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"todoList/config"
	"todoList/orm"
)

type UserClaims struct {
	orm.User
	jwt.StandardClaims
}

func GenToken(claims *UserClaims) (string, bool) {
	claims.ExpiresAt = time.Now().Add(config.Conf.TokenConfig.EffectTime).Unix()
	sign, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(config.Conf.TokenConfig.SecretKey)
	if err != nil {
		return "", false
	}
	return sign, true
}

func ParseToken(tokenString string) (*UserClaims, bool) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return config.Conf.TokenConfig.SecretKey, nil
	})
	if err != nil {
		return nil, false
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, false
	}
	return claims, true
}
