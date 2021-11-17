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
	reader := bufio.NewReader(os.Stdin)

	tableNum, _ := getInput("Creating a new bill (Please Enter the Table #): ", reader)
	b := newBill(tableNum)
	fmt.Println(" Bill Created for Table #", b.tableNum)

	return b
}

func promptOptions(b bill) {
	var total float64 = 0
	fmt.Print("\n")
	
	// Print current items
	for k, v := range b.items {
		var formatedItems string = fmt.Sprintf("%-15v ... $%v\n", k+":", v)
		fmt.Println(formatedItems)
		total += v
	}

	var formatedTip string = fmt.Sprintf("%-15v ... $%v\n", "Tip", b.tip)
	if b.tax != 0 {
		fmt.Println(formatedTip)
		total += b.tip
	}


	// Print current tax
	var formatedTax string = fmt.Sprintf("%-15v ... $%0.2f\n", "Tax", b.tax)
	if b.tax != 0 {
		fmt.Println(formatedTax)
		total += b.tax
	}


	if b.tax != 0 && total != 0 {
		fmt.Println("____________________________\n")
	}

	// Print current total
	var formatedTotal string = fmt.Sprintf("%-15v ... $%0.2f\n", "Total", total)
	if total != 0 {
		fmt.Println(formatedTotal)
	}

	if total != 0 {
		fmt.Print("\n")
	}

	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput(
		"Choose option (a - add item, s - save bill, t - add tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number...")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("item added -", name, "$"+price)
		promptOptions(b)
		
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number...")
			promptOptions(b)
		}
		b.updateTip(t)

		fmt.Println("tip has been updated to", tip)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("Bill for Table #" + b.tableNum + " has been saved", )
	default:
		fmt.Println("That was not a valid option...")
		promptOptions(b)
	}
}

func main() {

	mybill := createBill()
	promptOptions(mybill)

}