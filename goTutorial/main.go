package main

import (
	"fmt"
)

func main() {

	// for loop

	for i := 0; i < 5; i++ {
		fmt.Println("i value is", i)
	}

	/**********/
	
	val := 1
	val1 := 5

	for val < 10 {
		fmt.Println(val, " * ", val1, " = ", (int(val) * int(val1)))
		val++
	}

	/************/

	myArray := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}

	for i := 0; i < len(myArray); i++ {
		fmt.Println(i, "index array value is", myArray[i])
	}

	/***********/

	for index, value := range myArray {
		fmt.Printf("the value of index %v is %v\n", index, value)
	}

	/*********/

	for _, value := range myArray {
		fmt.Printf("the value is %v\n", value)
	}
}
