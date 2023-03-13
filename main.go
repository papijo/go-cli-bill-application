package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Get Input from Console Helper Function
func getInput( prompt string, r *bufio.Reader)(string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new Bill: ", reader)

	b:= newBill(name)
	fmt.Println("Created the bill -> ", b.name)

	return b

}


func promptOptions(b bill){
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Create option (a - add item, s - save, t - add tip): ", reader)
	
	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}

		//Add item to bill
		b.addItem(name, p)
		fmt.Println("Item added -> ", name, "$"+price)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}

		//Add item to bill
		b.updateTip(t)
		fmt.Println("Tip added -> ",  "$"+tip)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("You saved the file", b.name)
	default:
		fmt.Println("That was not a valid option")
		promptOptions(b)	
	}

}

func main() {
	myBill := createBill()
	promptOptions(myBill)
}