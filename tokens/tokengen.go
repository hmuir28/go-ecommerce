package tokengen

import (
	"os"
	"github.com/golang-jwt/jwt"
)

func GenerateToken(email, firstName, lastName string, userId string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": userId,
        "email": email,
		"firstName": firstName,
		"lastName": lastName,
    })

	secretKey := os.Getenv("SECRET_KEY")
	
	if secretKey == "" {
		secretKey = "thisismysecretkey"
	}

    tokenString, err := token.SignedString([]byte(secretKey))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}
