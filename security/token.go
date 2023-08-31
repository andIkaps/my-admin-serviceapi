package security

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenService interface {
	GenerateToken(userID string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

var SECRET_KEY = []byte(os.Getenv("JWT_KEY"))

func GenerateToken(userID string) (string, error) {
	claim := jwt.MapClaims{}
	claim["ID"] = userID
	claim["ExpiredAt"] = time.Now().Add(1 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signToken, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return signToken, err
	}

	return signToken, nil
}

func ValidateToken(token string) error {
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return err
	}

	return nil
}
