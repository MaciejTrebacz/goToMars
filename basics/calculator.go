package main

import "fmt"

func main() {
	var expenses, revenue, taxRate = GetInput()
	var earningsBeforeTax, earningsAfterTax, ratio = DoMath(expenses, revenue, taxRate)
	PrintResult(earningsBeforeTax, earningsAfterTax, ratio)
	formattedFV := fmt.Sprintf("Future value: %.1f\n", earningsAfterTax)
	fmt.Print(formattedFV)
	// multiline stringsS
	fmt.Printf(`1
						2
						 3`)
}

func PrintResult(earningsBeforeTax float64, earningsAfterTax float64, ratio float64) {
	fmt.Println("earnings before: ", earningsBeforeTax, "earnings after tax: ", earningsAfterTax, "ratio: ", ratio)
	fmt.Printf("variable: %v\n", earningsBeforeTax)
	fmt.Printf("variable formated 1 decimal places %.1f", earningsAfterTax)
}

func GetInput() (expenses, revenue, taxRate float64) {
	fmt.Print("Input earning before tax: ")
	fmt.Scan(&revenue)

	fmt.Print("Input expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Input tax rate: ")
	fmt.Scan(&taxRate)
	return
}

func DoMath(expenses, revenue, taxRate float64) (EBT, EAT, ratio float64) {
	EBT = revenue - expenses
	EAT = EBT - (EBT * (taxRate / 100))
	ratio = EBT / EAT
	return
}
