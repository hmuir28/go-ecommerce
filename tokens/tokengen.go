package tokengen

import (
	"os"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	User_ID string
	Email string
	First_Name string
	Last_Name string
	jwt.RegisteredClaims
}

func GenerateToken(email, firstName, lastName string, userId string) (string, string, error) {
	// Define the expiration time for the access token and refresh token
	accessTokenExpiry := time.Now().Add(15 * time.Minute) // Access token valid for 15 minutes
	refreshTokenExpiry := time.Now().Add(24 * time.Hour)  // Refresh token valid for 24 hours

	// Create the JWT claims, which includes the username and the expiry time
	claims := &Claims{
		User_ID: userId,
		Email: email,
		First_Name: firstName,
		Last_Name: lastName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessTokenExpiry),
		},
	}


	jwtKey := os.Getenv("SECRET_KEY")
	
	if jwtKey == "" {
		jwtKey = "thisismysecretkey"
	}

	
	// Generate the access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenString, err := accessToken.SignedString(jwtKey)
	if err != nil {
		return "", "", err
	}

	// Generate the refresh token (typically without expiration info in claims)
	refreshClaims := &Claims{
		User_ID: userId,
		Email: email,
		First_Name: firstName,
		Last_Name: lastName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshTokenExpiry),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(jwtKey)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func UpdateAllTokens(token, refreshToken, userId string) {

}
