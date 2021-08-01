package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {

	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {

	// user input, switch statement

	reader := bufio.NewReader((os.Stdin)) // bufio.NewReader used to read user info
	// os.Stdin is to receiver typed info from console

	// fmt.Print("Create a new bill, entre name:")	// show message to enter name for the bill
	// name, _ := reader.ReadString('\n')			// get user input and enter key
	// name = strings.TrimSpace(name)				// trim any empty space in the user input

	name, _ := getInput("Create a new bill, entre name: ", reader)

	b := newBill(name) // use the entered value to prepare a bill
	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOptions(b bill) {

	reader := bufio.NewReader(os.Stdin) // bufio.NewReader used to read user info
	// os.Stdin is to receiver typed info from console

	options, _ := getInput("Choose option (a - add item, s - save bill, t - add tip, q - quit \n", reader) // get user option

	switch options {
	case "a":
		name, _ := getInput("Item name:", reader)
		price, _ := getInput("Item price: ", reader)
		p, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("Invalid character")
			promptOptions(b)
		} else {
			b.addItem(name,p)
			fmt.Printf("You chose %v and price is $%0.2f \n", name, p)
			promptOptions(b)
		}

	case "s":

		b.save()
		fmt.Println("You chose s. Bill saved successfully \n", b.name)

	case "t":
		tip, _ := getInput("Enter tip amount: ", reader)
		t, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("Invalid character")
			promptOptions(b)
		} else {

			b.updateTip(t)
			fmt.Printf("Your tip added $%0.2f \n", t)
			promptOptions(b)
		}

	case "q":
		fmt.Println("You chose q, stopping execution")
		break

	default:
		promptOptions(b)
		fmt.Println(options)
	}
}

func main() {

	// reader --- https://www.youtube.com/watch?v=20HlPtQuc_g

	mybill := createBill() // create bill
	promptOptions(mybill)  // do user option

	// fmt.Println(mybill)
}
