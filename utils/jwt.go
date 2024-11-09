package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/goschool/crud/types"
)

var secretKey = []byte("my-super-secret-secret")

func CreateToken(user types.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":    user.Name,
		"id":      user.ID,
		"email":   user.Email,
		"expires": time.Now().Add(time.Hour * 24).Unix(),
	}, nil)

	tokentString, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("Failed to sign token: %w", err)
	}

	return tokentString, nil
}
