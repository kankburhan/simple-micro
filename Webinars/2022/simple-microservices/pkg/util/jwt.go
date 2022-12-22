package util

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"

	//"golang.org/x/oauth2/jwt"
	"time"
)

type ClaimsJWT struct {
	jwt.StandardClaims
	UserID string `json:"user_id"`
	Email  string `json:"email"`
}

func GenerateJwt(
	userid string,
	email string,
	issuer string,
	secret string,
	expireHour int,
) (string, error) {
	claims := ClaimsJWT{
		StandardClaims: jwt.StandardClaims{
			Issuer:    issuer,
			ExpiresAt: time.Now().Add(time.Duration(expireHour) * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserID: userid,
		Email:  email,
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateJwt(token string, secret string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})
}
