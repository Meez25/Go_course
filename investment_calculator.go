package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var (
		investmentAmount   float64
		expectedReturnRate float64
		years              float64
	)

	fmt.Print("Enter your investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the expected return rate of your investment: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter the number of years you are waiting to collect: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Printf("Future Value: %.2f\n", futureValue)
	fmt.Printf("Future Value (adjusted for Inflation): %.2f", futureRealValue)
}
