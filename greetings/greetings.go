package greetings

import "fmt"

// Hello :
//
//	a function whose name starts with a capital letter can be called by a function
//	not in the same package
func Hello(name string) string {
	//the := operator is a shortcut for declaring and initializing a variable in one line
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}