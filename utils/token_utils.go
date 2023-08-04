package utils

import (
	"time"

	"github.com/4epyx/authrpc/pb"
	"github.com/golang-jwt/jwt"
)

func GenerateUserAccessToken(user *pb.User, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, jwt.MapClaims{
		"user_id":    user.Id,
		"user_email": user.Email,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(secretKey)
}

func GetJWTClaims(token, secretKey string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}

	if _, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	}); err != nil {
		return nil, err
	}

	return claims, nil
}
