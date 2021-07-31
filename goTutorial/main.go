package main

import (
	"fmt"
)

func updateName(x string) string {

	fmt.Println("x is", x)			//	x is ka

	x = "pa"

	fmt.Println("x is", x)			// x is pa

	return x
}

func updateMenu(y map[string]float64) {

	y["coffee"] = 5.99

	fmt.Println("y before update",y)	// y before update map[coffee:5.99 tea:0.99 water:0.49]

	y["coffee"] = 4.99

	fmt.Println("y after update",y)		// y after update map[coffee:4.99 tea:0.99 water:0.49]
}

func main() {

	// Pass By Value ! IMPORTANT *** how data is handled in group a and group b types. THERE IS DIFFERENC IN GROUP A AND B

	// Learn this in https://www.youtube.com/watch?v=LBLN4Wu5U4w

	// Group A types = strings, ints, booleans, floats, arrays and struct

	name := "ka"

	fmt.Println("name is", name)	// name is ka

	var newName = updateName(name) 

	fmt.Println("name is", name)	// name is ka
	fmt.Println("newName is", newName)	// newName is pa

	// Group B types = slice, maps and functions

	menu := map[string]float64{
		"water": 0.49,
		"tea":   0.99,
	}

	updateMenu((menu))

	fmt.Println(menu)	// map[coffee:4.99 tea:0.99 water:0.49]

}
