package main

import (
	"fmt"
)

type bill struct {
	name string
	items map[string]float64
	tip float64
}

// create a new bill

func newBill(name string) bill {
	b := bill {
		name:name,
		items: map[string]float64{"cake" : 3.99, "water" : 0.49},
		tip:0,
	}
	return b
}

// formatting the bill

func (b bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	// list items
	for key, value := range b.items {
		fs += fmt.Sprintf("%-10v ...$%v \n", key+":", value)	// -10 will reserves 10 characters long space for each keys, the prints value  example cake:      ...$3.99
		total += value
	}

	// total
	fs += fmt.Sprintf("%-10v ...$%0.2f", "total:", total)	// -10 will reserves 10 characters long space for each keys, the prints value  example total:       ...$4.48
	
	return fs
}