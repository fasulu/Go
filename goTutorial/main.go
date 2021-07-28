package main

import (
	"fmt"
)

func main() {

	// array	fixed with number of arrays cant be appended

	var name1 [4]string = [4]string{"a", "b", "c", "d"} // method one
	var name2 = [4]string{"e", "r", "y", "u"}       	// method 2
	name3 := [4]int{2, 8, 9, 10}           				// method 3 


	fmt.Println("hello", name1, name2, name3)

	name3[2] = 10000		// changes 3 array value 9 to 10000

	fmt.Println(name3)

	fmt.Println(name1, len(name1))	// len gives length of the array

	// slice (use arrays under the hood) In slices arrays can be appened

	var valuesofteam = []int{5,10,8,9}		// empty [] signifies it is slice

	fmt.Println(valuesofteam)

	valuesofteam[1] = 9

	fmt.Println(valuesofteam)

	valuesofteam = append(valuesofteam, 10)		// values can be appened in slices

	fmt.Println(valuesofteam)

	// slice ranges

	rangeOne := name1[1:3]	// picks array 1 and 2 excluding 3 result is [b c]

	rangeTwo := name1[2:]	// picks array 2 upto end of the array result is [c d]

	rangeThree := name1[:2]	// picks from start of an array to 1st value of array excluding 2 result is [a b]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

}
