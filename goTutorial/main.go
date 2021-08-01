package main

import (
	"fmt"
)

func main() {
	
	// Struct --- https://www.youtube.com/watch?v=zfMhSFpBSaw
	// Struct means structure. Is used for custom build structure for data object 
	// ie blue print of data stucture which is equivalent to class objects in other object oriented languages
	// Here we make a custom build data structure called bill 
	// Using Receiver Functions

	mybill := newBill("ray's bill")

	fmt.Println(mybill)	// creates a new bill using the bill.go. result is {ray's bill map[] 0}

	mybill.addItem("cake", 3.99)
	mybill.addItem("water", 0.49)
	mybill.addItem("coffee", 1.99)
	mybill.addItem("hot choclat", 2.49)

	mybill.updateTip(10)
	
	fmt.Println(mybill.format())
}
