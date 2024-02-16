package main

import "fmt"

func main() {
	var expenses float64
	var revenue float64
	var taxRate float64

	fmt.Print("Input earning before tax: ")
	fmt.Scan(&revenue)

	fmt.Print("Input expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Input tax rate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses
	earningsAfterTax := earningsBeforeTax - (earningsBeforeTax * (taxRate / 100))
	ratio := earningsBeforeTax / earningsAfterTax

	fmt.Println("earnings before: ", earningsBeforeTax, "earnings after tax: ", earningsAfterTax, "ratio: ", ratio)
	fmt.Printf("variable: %v\n", earningsBeforeTax)
	fmt.Printf("variable formated 1 decimal places %.1f", earningsAfterTax)

	// formatting strings
	formattedFV := fmt.Sprintf("Future value: %.1f\n", earningsAfterTax)
	fmt.Print(formattedFV)

	// multiline stringsS
	fmt.Printf(`1
						2
						 3`)
}
