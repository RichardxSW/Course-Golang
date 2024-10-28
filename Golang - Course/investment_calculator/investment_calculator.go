package main

import (
	"fmt"
	"math"
)

	const inflationRate = 2.5
func main() {
	// var investmentAmount float64 = 1000
	// var expectedReturnRate = 5.5
	// var years float64= 10
	// var investmentAmount, years float64 = 1000, 10
	// expectedReturnRate := 5.5
	var investmentAmount, expectedReturnRate, years float64

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Enter years: ")
	fmt.Scan(&years)

	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future real value: %.2f\n", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (FV, RFV float64) {
	FV = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	RFV = FV / math.Pow(1+inflationRate/100, years)
	// return FV, RFV
	return
}