package user

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtkey = []byte("taadis_secret")

type Claims struct {
	UserId string
	jwt.StandardClaims
}

// GenerateToken
func GenereateToken(userId string) (string, error) {
	claims := &Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "taadis.com",
			Subject:   "json_web_token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})
	return token, Claims, err
}
