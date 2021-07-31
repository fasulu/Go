package main

import (
	"fmt"
)

func updateName(x string) string {

	fmt.Println("x is", x)			//	x is ka

	x = "pa"

	fmt.Println("x is changed to pa: ", x)			// x is pa

	return x
}

func updateNameUsingPointer(y *string) {

	fmt.Println("y in func updateNameUsingPointer is ", *y)		// result is y in func updateNameUsingPointer is  ka

	*y = "ma"

	fmt.Println("y in func updateNameUsingPointer after updated is : ", *y)		// result is y in func updateNameUsingPointer after updated is : ma


}

func main() {
	
	// POINTER - How pointers work in Go --- https://www.youtube.com/watch?v=4B2rwYvuiBo
	
	// Pointer is used to point the data stored memory location

	name := "ka"

	fmt.Println("memory address of name is", &name)	// &name gives memory address, &name is called pointer
	
	mem := &name
	
	fmt.Println("mem is stroed the memory addres", mem)	// printers memory address stored in mem, ie pointer 

	fmt.Println("*mem shows the value stored in its memory address", *mem)		// *mem prints the value stroed in that memory address

	updateName(name) 

	fmt.Println()

	// how to update value using pointer

	fmt.Println("name before update :", name)	// result is name before update : ka
	
	updateNameUsingPointer(mem)		// passing memory location to a function updateNameUsingPointer
	
	fmt.Println("name after update :", name)	// result is name after update : ma
}
