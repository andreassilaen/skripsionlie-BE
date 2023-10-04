package skirpsionlineBE

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func (s Service) TokenUser(ctx context.Context) error {
	// Seed the random number generator with the current time
	var (
		err	error
	)
    // Seed the random number generator with the current time
    rand.Seed(time.Now().UnixNano())

    // Specify the length of the random string
    length := 10

    // Generate a random string
    randomString := generateRandomString(length)
    fmt.Println("Random String:", randomString)

	return err
}

func generateRandomString(length int) string {
    // Define the character set from which to choose
    charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

    // Create a byte slice to store the random string
    randomBytes := make([]byte, length)

    // Populate the byte slice with random characters
    for i := 0; i < length; i++ {
        randomBytes[i] = charSet[rand.Intn(len(charSet))]
    }

    // Convert the byte slice to a string and return it
    return string(randomBytes)
}