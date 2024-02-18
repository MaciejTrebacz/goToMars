package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64   // doesn't have initial value as we getting input from user
	var expectedReturnRate float64 //we can use short type assignment if we don't need inferred type
	var years float64

	outputText("Provide investment amount")
	fmt.Scan(&investmentAmount) // pointer

	outputText("Provide years to invest: ")
	fmt.Scan(&years) // pointer

	outputText("Provide return rate: ")
	fmt.Scan(&expectedReturnRate) // pointer

	var firstVariable, secondVariable float64 = 1, 2 // we can declare same type variables in one line
	thirdVariable, forthVariable := 1.0, 2.0

	fmt.Println(firstVariable, secondVariable, thirdVariable, forthVariable)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

	receiveAndReturn("1", "2")
}

func outputText(text string) {
	fmt.Print(text)
}

func receiveAndReturn(variable1, variable2 string) (string, string) { // we can return by just writing name of variable like "a string"
	fmt.Printf(variable1, variable2)
	return "a", "b"
}
