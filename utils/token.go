package utils

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	data interface{}
	jwt.StandardClaims
}

// var expr int64 = 5

const (
	sercet_Key = "x14n_key"
)

func GetToken(key string, exprTime int) (string, error) {
	expriationTime := time.Now().Add(5 * time.Minute)

	//生成payload
	claims := &Claims{
		data: key,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expriationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(sercet_Key)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func CheckToken(token string) bool {

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return sercet_Key, nil
	})
	if err != nil {
		log.Println("token解析错误", err)
		return false
	}
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		log.Println("token过期了")
		return false
	}
	if !tkn.Valid {
		log.Println("token不正确")
		return false
	}
	return true
}
