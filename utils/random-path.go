package utils

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/aidarkhanov/nanoid/v2"
	"github.com/reisvitt/url-shortener-go/configs"
)

func GerenateRandomPath() (string, error) {

	rand.New(rand.NewSource(time.Now().UnixNano()))
	// Generate a number between 5 and 10
	min := 5
	max := 10
	randomNumber := rand.Intn(max-min+1) + min

	// Generate a random alphanumeric string of the desired length
	shortCode, err := nanoid.GenerateString(configs.ALPHANUMERIC, randomNumber)
	if err != nil {
		fmt.Println("Error generating short code:", err)
		return "", err
	}

	return shortCode, nil
}
