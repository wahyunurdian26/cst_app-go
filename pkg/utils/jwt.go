package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var jwtSecret = []byte("XYYZ")

// GenerateJWT membuat token JWT untuk user
func GenerateJWT(userID uuid.UUID, email string, role string, group string, subgroup string, digital string) (string, error) {
	claims := jwt.MapClaims{
		"sub":      userID.String(),
		"email":    email,
		"role":     role,
		"group":    group,
		"subgroup": subgroup,
		"digital":  digital,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token berlaku 24 jam
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
