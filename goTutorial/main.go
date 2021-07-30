package main

import (
	"fmt"
)

var score = 19.2	// to use this score variable, it needs to be declared outside function

func main() {

	// package scope, using variables declared in different go packages

	sayHello("ray")

	for _, v := range names {
		fmt.Println((v))
	}

	showScore()
}
