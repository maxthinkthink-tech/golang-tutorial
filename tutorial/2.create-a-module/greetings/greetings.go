package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello :
//
//	a function whose name starts with a capital letter can be called by a function
//	not in the same package
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	//the := operator is a shortcut for declaring and initializing a variable in one line
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		msg, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = msg
	}
	return messages, nil
}

// randomFormat starts with a lowercase letter,
// making it accessible only to code in its own package (in other words, it's not exported).
func randomFormat() string {
	// declare and initialize a slice
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
