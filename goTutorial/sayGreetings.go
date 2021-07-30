package main

import "fmt"

// to use this names array, it needs to be declared outside function
var names = []string{"ka", "ma", "cha", "da", "tha", "pa", "ra", "na", "va"}

func sayHello(n string) {

	fmt.Printf("hello %v \n", n)
}

func showScore() {
	fmt.Println("your score is", score)
}