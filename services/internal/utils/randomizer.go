package utils

import (
	"time"
	"math/rand"
)


func GenerateRandomString(length int) string {
	// Define the characters that can be used in the random string
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Create a byte slice to store the random string
	randomString := make([]byte, length)

	// Populate the byte slice with random characters
	for i := 0; i < length; i++ {
		randomString[i] = charset[rand.Intn(len(charset))]
	}

	return string(randomString)
}