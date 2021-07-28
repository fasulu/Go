package main

import (
	"fmt"
)

func main() {

	// string

	var name1 string = "paris" // method one
	var name2 = "london"       // method 2
	var name3 string           // method 3 withou value
	name4 := "tokyo"           // method 4 most popular in go

	fmt.Println("hello", name1, name2, name3, name4)

	// uint int int8 int16 int32 int64	--- each of these has specific value range. verify golang docs https://golang.org/ref/spec#Numeric_types

	var val1 int = 18 // method one
	var val2 = 19     // method 2
	var val3 int16    // method 3 withou value
	val4 := 20        // method 4 most popular in go. if stored value string treats as string, if stored value is int it treats as int

	fmt.Println("hello", val1, val2, val3, val4, val1+val2)

	// float32 float64

	var flt1 float32 = 13.89
	var flt2 float64 = 4335353455343463534534534.34534
	flt3 := 465465.57456
	flt4 := (float64(flt1) + float64(flt2))				// method of adding 2 floats

	fmt.Println(flt1, flt2, "n makes to print the following in newline \n ", flt3)
	fmt.Println(flt4)

	// formated string printing. Learn more from https://pkg.go.dev/fmt#pkg-index

	fmt.Printf("\n%v plus %v is equal to %v \n", val1, val2, int(val1) + int(val2))

	fmt.Printf("hello %v this will make a qoute on 2nd arguement name2 %q ", name1, name2)

	fmt.Printf("what is the type of arguement %T \n", name4)	// "london"
	
	fmt.Printf("this will give float numbers %0.1f \n", flt1) // 13.9



}
