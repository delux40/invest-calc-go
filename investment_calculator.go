package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formatedFV := fmt.Sprintf("Future value: %.4f \n", futureValue)
	fmt.Print(formatedFV)
	formatedFRV := fmt.Sprintf("Future real value: %.4f", futureRealValue)
	fmt.Print(formatedFRV)
}
