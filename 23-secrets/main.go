package main

import (
	"encoding/base32"
	"fmt"
)

func main() {
	// Define a 16-character
	secret := "my_secret_key_01"

	// Convert the string to bytes
	myBytes := []byte(secret)

	// Encode the bytes using Base32
	myEncodedBytes := base32.StdEncoding.EncodeToString(myBytes)

	fmt.Println(myEncodedBytes)
	fmt.Println(secret)
}
