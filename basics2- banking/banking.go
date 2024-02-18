package main

import "fmt"

func main() {
	var accountBalance float64 = 1000

	fmt.Println("Welcome to GO BANK...")
	fmt.Println("What you want to do...")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Scan(&choice) // number user enters

	if choice == 1 {
		fmt.Println("Your balance is: ", accountBalance)
	}

	fmt.Println("Your choice: ", choice)

}
