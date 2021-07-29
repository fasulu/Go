package main

import (
	"fmt"
)

func main() {

	// Booleans, if Conditions, continue and break

	males := 25
	females := 15

	fmt.Println(males <= females) //false
	fmt.Println(females >= males) //false
	fmt.Println(males > females)  //true
	fmt.Println(females != males) //true

	if males > 25 {

		fmt.Println("males are not more than 25")

	} else if females > males {

		fmt.Println("females are less than males")

	} else {

		fmt.Println("there are diffence in female and male counts")

	}

	// run the following code and learn how continue and break works

	myArray := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}

	for index, value := range myArray {

		if index == 1 {

			fmt.Println("continuing at position", index)
			continue
		}

		if index > 2 {
			fmt.Println("breaking at position", index)
			break
		}

		fmt.Printf("the value at position %v is %v \n", index, value)
	}

}
