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

	outputText("Enter your investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the expected return rate of your investment: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter the number of years you are waiting to collect: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years, inflationRate)

	formattedFV := fmt.Sprintf("Future Value: %.1f", futureValue)
	formattedRFV := fmt.Sprintf("Future Value with inflation: %.1f", futureRealValue)

	fmt.Print(formattedFV)
	fmt.Print(formattedRFV)

}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount float64, expectedReturnRate float64, years float64, inflationRate float64) (float64, float64) {
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	return futureValue, futureRealValue
}
