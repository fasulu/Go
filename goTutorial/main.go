package main

import (
	"fmt"
	"math"
	"strings"
)

func sayGreeting(names string) {
	fmt.Printf("Good morning %v \n", names)
}

func sayBye(names string) {
	fmt.Printf("Goodbye %v \n", names)
}

func cycleNames(names []string) {
	for _, value := range names {
		fmt.Printf("Name is %v \n", value)
	}
}

// In below function the first float64 is what type we are passing for calculation,
// the second float64 is what type of value we are going to return

func circleArea(radius float64) float64 {
	return math.Pi * radius * radius
}

func recycleNames(names []string, funt func(string)) {
	for _, value := range names {
		funt(value)
	}
}

//* to return multiple values specify that in type like string/int/float etc
func getIntials(n string) (string, string) {

	s := strings.ToUpper(n)			// here strings.ToUpper is syntax to change n values to uppercase 

	fmt.Println("s is", s)			// result is "s is RAY CAT"

	list := strings.Split(s, " ")	// split the words into array list

	fmt.Println("list is", list)	// result is array "list is [RAY CAT]"

	var initials []string			// create empty array as initials
	for _, v := range list {
		initials = append(initials, v[:1])
	}

	fmt.Println("initials is", initials)	// result array initials is "initials is [R C]"

	if len(initials) > 1 {
		return initials[0], initials[1]		// check if more than one to return
	}
	return initials[0], "_"					// if not send one initial and one empty as "_"

}

func main() {

	// functions and returns

	myArray := []string{"ka", "ma", "cha", "da", "tha", "pa", "ra", "na", "va"}

	sayGreeting("ray")
	sayBye("chat")

	cycleNames([]string{"ray", "cat"})

	cycleNames(myArray)

	recycleNames(myArray, sayGreeting)

	areaResult1 := circleArea(12.5)
	areaResult2 := circleArea(9)

	fmt.Printf("Area Result is %v \n", areaResult1)
	fmt.Printf("Area Result is %v \n", areaResult2)
	fmt.Printf("Formatted string for less decimal point is %0.2f \n", areaResult2)

	// functions with multiple returns

	fn1, sn1 := getIntials("ray cat")
	fmt.Printf("initials of list is %v, %v \n",fn1, sn1)

	fn2, sn2 := getIntials("sha")
	fmt.Printf("initial of list is %v, %v \n", fn2, sn2)
}
