package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// createJWT creates a new JWT token with the given claims and secret key.
func createJWT(claims jwt.MapClaims, secretKey []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// readJWT reads and validates a JWT token with the given token string and secret key.
func readJWT(tokenString string, secretKey []byte) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("Failed to extract claims from token")
	}
}

func main() {
	fmt.Println("***JWT Fun")
	// Define the secret key for signing the token. This should be kept secret and not hard-coded in production.
	secretKey := []byte("my_secret_key")
	fmt.Println("secret key (byte array) = ", secretKey)
	str1 := string(secretKey[:])
	fmt.Println("Back to String = ", str1)

	// Create a new JWT token with claims
	claims := jwt.MapClaims{
		"sub": "1234567890",                         // Subject
		"iss": "my_issuer",                          // Issuer
		"exp": time.Now().Add(time.Hour * 1).Unix(), // Expiration time (1 hour from now)
	}

	// Call createJWT to create the token
	tokenString, err := createJWT(claims, secretKey)
	if err != nil {
		fmt.Println("Failed to create token:", err)
		return
	}
	fmt.Println("JWT Token:", tokenString)

	// Call readJWT to read and validate the token
	claims, err = readJWT(tokenString, secretKey)
	if err != nil {
		fmt.Println("Failed to read token:", err)
		return
	}

	// Extract claims from the parsed token
	subject := claims["sub"].(string)
	issuer := claims["iss"].(string)
	expiration := time.Unix(int64(claims["exp"].(float64)), 0)

	fmt.Println("Subject:", subject)
	fmt.Println("Issuer:", issuer)
	fmt.Println("Expiration:", expiration)
}
