package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5      // cannot change
	var investmentAmount float64   // doesn't have initial value as we getting input from user
	var expectedReturnRate float64 //we can use short type assignment if we don't need inferred type
	var years float64

	fmt.Print("Provide investment amount: ")
	fmt.Scan(&investmentAmount) // pointer

	fmt.Print("Provide years to invest: ")
	fmt.Scan(&years) // pointer

	fmt.Print("Provide return rate: ")
	fmt.Scan(&expectedReturnRate) // pointer

	var firstVariable, secondVariable float64 = 1, 2 // we can declare same type variables in one line
	thirdVariable, forthVariable := 1.0, 2.0

	fmt.Println(firstVariable, secondVariable, thirdVariable, forthVariable)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

}
