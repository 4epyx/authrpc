package util

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// TODO: put away []byte conversion

func GenerateUserAccessToken(user *User, secretKey []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    user.Id,
		"user_email": user.Email,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(secretKey)
}

func GetJWTClaims(token string, secretKey []byte) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}

	if _, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	}); err != nil {
		return nil, err
	}

	return claims, nil
}
