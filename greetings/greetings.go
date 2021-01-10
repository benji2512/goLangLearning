package greetings

import (
	"errors"
	"fmt"
)

// Hello will return a greeting for the named person
func Hello(name string) (string, error) {
	// Return an error if name is not supplied
	if name == "" {
		return "", errors.New("empty name")
	}

	// If name was received, return a greeting that embeds the name in a message
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
