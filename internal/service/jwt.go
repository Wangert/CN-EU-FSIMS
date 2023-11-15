package service

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang/glog"
	"github.com/spf13/viper"
	"time"
)

const JWTSECRET = "FSIMSJWT666"

type FsimsUserClaims struct {
	UUID string
	jwt.StandardClaims
}

// CreateJwtToken Create JWT Token
func CreateJwtToken(uuid string) (string, error) {
	// 2 hours
	var TokenExpireDuration = time.Second * time.Duration(viper.GetInt("jwt.timeout")) * 2

	glog.Info("duration:", TokenExpireDuration)
	//jwt secert
	jwtSecret := []byte(JWTSECRET)

	c := FsimsUserClaims{
		uuid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "FSIMS-SYSTEM",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString(jwtSecret)
}

// ParseToken parse jwt token
func ParseToken(tokenString string) (*FsimsUserClaims, error) {
	// jwt secret
	var jwtSecret = []byte(JWTSECRET)
	// parse token
	token, err := jwt.ParseWithClaims(tokenString, &FsimsUserClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
	})

	if err != nil {
		glog.Errorln("parse error")
		return nil, err
	}

	fmt.Println("claims:")
	fmt.Println(token.Claims.(*FsimsUserClaims))

	if claims, ok := token.Claims.(*FsimsUserClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
