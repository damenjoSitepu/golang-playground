package main

import (
	"fmt"
)

// Say Hello Function
func sayHello(name string, year int8) string {
	var message string = fmt.Sprintf("Hi, my name is %v and now i am %v years old!", name, year)
	return message
}

// Main Function
func main() {
	fmt.Println(sayHello("Damenjo", 23))
}
