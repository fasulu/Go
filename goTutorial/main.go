package main

import (
	"fmt"
)

func main() {
	
	// Struct --- https://www.youtube.com/watch?v=zfMhSFpBSaw
	// Struct means structure. Is used for custom build structure for data object 
	// ie blue print of data stucture which is equivalent to class objects in other object oriented languages
	// Here we make a custom build data structure called bill 

	mybill := newBill("ray's bill")

	fmt.Println(mybill)	// creates a new bill using the bill.go. result is {ray's bill map[] 0}
	
	fmt.Println(mybill.format())
}
