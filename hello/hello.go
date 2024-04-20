// In Go, code executed as an application must be in a main package.
package main

import "fmt"
import "example.com/greetings"

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
