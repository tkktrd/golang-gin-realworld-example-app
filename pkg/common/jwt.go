package common

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
	"tkktrd/golang-gin-realworld-example-app/config"
)

const (
	userIdKey = "user_id"
	emailKey = "email"
	iatKey = "iat"
	expKey = "exp"
)

type JwtClaims struct {
	ID string `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJwtToken(id uint, email string, config *config.Config) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		userIdKey: id,
		emailKey: email,
		iatKey: now.Unix(),
		expKey: now.Add(time.Minute * 60).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.Server.JwtSecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}