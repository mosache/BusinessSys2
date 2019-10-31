package service

import (
	"encoding/base64"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

const (
	secretKey = "dhdiadhusnnvnvaopdpeiejpas"
)

type TokenData struct {
	UserId int
	jwt.StandardClaims
}

func GetNewToken(userId int) string {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodES256, &TokenData{UserId: userId})

	result, err := jwtToken.SignedString(secretKey)

	if err != nil {
		return ""
	}

	result = base64.StdEncoding.EncodeToString([]byte(result))

	return result
}

func CheckToken(token string) (data interface{}, err error) {
	var bytes []byte
	bytes, err = base64.StdEncoding.DecodeString(token)

	parsedToken, parseErr := jwt.ParseWithClaims(string(bytes), &TokenData{}, func(token *jwt.Token) (i interface{}, e error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexcepted signing method:%v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if claims, ok := parsedToken.Claims.(*TokenData); ok && parsedToken.Valid {
		data = claims
		fmt.Printf("%v %v", claims.UserId, claims.StandardClaims.ExpiresAt)
	} else {
		fmt.Println(err.Error())
		err = parseErr
		return
	}

	return
}
