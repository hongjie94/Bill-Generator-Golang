package main

import (
	"fmt"
	"os"
)

type bill struct {
	tableNum  string
	items map[string]float64
	tip	float64
	tax float64
	taxRate float64
}

// make new bills
func newBill(tableNum string) bill {
	b := bill{
		tableNum:  tableNum,
		items: map[string]float64{},
		tip: 0,
		tax: 0,
		taxRate: 0.08,
	}
	return b
}

// add item to bill
func (b *bill) addItem(tableNum string, price float64) {
	b.items[tableNum] = price
	b.tax = price*b.taxRate
}

// format the bill
func (b *bill) format() string {
	fs := "Bill breakdown:\n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v\n", k+":", v)
		total += v
	}

	// add tip
	fs += fmt.Sprintf("%-25v ...$%v\n", "tip:", b.tip)

	// add tax
	fs += fmt.Sprintf("%-25v ...$%0.2f\n", "tax:", b.tax+(total*b.taxRate))

	// add total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total+(b.tip + b.tax))

	return fs
}

// update tip
func (b *bill) updateTip(tip float64) {
	(*b).tip = tip
}

// save bill
func (b *bill) save() {
	data := []byte(b.format())
	err := os.WriteFile("bills/"+"Table #"+b.tableNum+".txt", data, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Bill saved to file")
}