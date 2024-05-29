package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("What is your revenue ? ")
	fmt.Scan(&revenue)

	fmt.Print("What is your expenses ? ")
	fmt.Scan(&expenses)

	fmt.Print("What is your tax rate ? ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses
	earningsAfterTax := earningsBeforeTax * (1 - taxRate/100)
	ratio := earningsBeforeTax / earningsAfterTax

	fmt.Printf("EBT : %f$, profit: %f$, ratio: %f\n", earningsBeforeTax, earningsAfterTax, ratio)
}
