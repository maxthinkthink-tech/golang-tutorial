// In Go, code executed as an application must be in a main package.
package main

import (
	"fmt"
	"log"
	// go mod edit -replace example.com/greetings=../greetings
	"example.com/greetings"
)

func testHello(name string) {
	log.SetPrefix("testHello: ")
	log.SetFlags(0)
	message, err := greetings.Hello(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}

func testHellos(name []string) {
	log.SetPrefix("testHellos: ")
	log.SetFlags(0)
	messages, err := greetings.Hellos(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}

func main() {
	//fmt.Println("-----------")
	//testHello("alan")
	//fmt.Println("-----------")
	//testHello("")

	fmt.Println("-----------")
	testHellos([]string{"a", "b"})
}
