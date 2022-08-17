package utils

import (
	"math/rand"
	"strings"
	"time"
)

func GetRandomText() string {
	rand.Seed(time.Now().Unix())
	var output strings.Builder

	charSet := "abcdedfghijklmnopqrstABCDEFGHIJKLMNOP"
	length := 20

	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}

	return output.String()
}
