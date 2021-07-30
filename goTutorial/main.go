package main

import (
	"fmt"
)

func main() {

	// Maps - used to store values in array with keys and values

	menu := map[string]float64 {
		
		"soup" : 1.99,
		"salad" : 0.99,
		"water" : 0.49,
		"coffee" : 0.99,
	}

	fmt.Println(menu["soup"])	// gives 1.99
	fmt.Println(menu)			// gives map[coffee:0.99 salad:0.99 soup:1.99 water:0.49]

	//

	phonelist := map[int]string {
		54664654: "ray",
		98765754: "sha",
		23541845: "bee",
	}

	fmt.Println(phonelist[54664654])
	fmt.Println(phonelist)

	// looping maps

	for k,v := range phonelist {
		fmt.Println(k, "-", v)			// result will be 54664654 - ray
														//98765754 - sha
														//23541845 - bee
	}

	// updating the map

	phonelist[54664654] = "rayh"	// change value of 54664654 to rayh

	fmt.Println(phonelist)		//result is map[23541845:bee 54664654:rayh 98765754:sha]
}
