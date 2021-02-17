package common

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"qi.xiao/ginessential/model"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

//登录成功后发放token
func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "qixiao",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	//从tokenString中解析出claim返回
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}

//"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjYsImV4cCI6MTYxNDE2NDg3MywiaWF0IjoxNjEzNTYwMDczLCJpc3MiOiJxaXhpYW8iLCJzdWIiOiJ1c2VyIHRva2VuIn0.Q1ebufQtByMSxdCeXXyw0xMNPLjL4tI4mzEq3jX7114"
//1.协议头 2.20-27行信息 3.前两部分加key组成的hash值
