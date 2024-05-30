package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	revenue, err := askForUserInput("What is your revenue ? ")

	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err = askForUserInput("What is your expenses ? ")

	if err != nil {
		fmt.Println(err)
		return
	}
	taxRate, err = askForUserInput("What is your tax rate ? ")

	if err != nil {
		fmt.Println(err)
		return
	}

	earningsBeforeTax, earningsAfterTax, ratio := calculateEarnings(revenue, expenses, taxRate)

	saveCalculationToFile(earningsBeforeTax, earningsAfterTax, ratio)

	fmt.Printf("EBT : %f$, profit: %f$, ratio: %f\n", earningsBeforeTax, earningsAfterTax, ratio)
}

func askForUserInput(text string) (float64, error) {
	var value float64
	fmt.Print(text)
	fmt.Scan(&value)
	if value <= 0 {
		return 0, errors.New("The value must be greater than 0")
	}
	return value, nil
}

func calculateEarnings(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := revenue - expenses
	earningsAfterTax := earningsBeforeTax * (1 - taxRate/100)
	ratio := earningsBeforeTax / earningsAfterTax
	return earningsBeforeTax, earningsAfterTax, ratio
}

func saveCalculationToFile(ebt float64, eat float64, ratio float64) {
	// valuesAsString := fmt.Sprint(ebt, eat, ratio)
	valuesAsString := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.1f\n", ebt, eat, ratio)
	err := os.WriteFile("calculations.txt", []byte(valuesAsString), 0644)
	if err != nil {
		fmt.Println("Could not write to file")
	}
}
