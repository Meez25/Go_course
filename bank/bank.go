package main

import (
	"example.com/bank/fileops"
	"fmt"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {

	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("--------------")
		// panic(err)
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())
	for {
		presentOptions()

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Println("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greated than 0.")
				continue
			}
			accountBalance += depositAmount
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			fmt.Println("The new balance is", accountBalance)
		case 3:
			fmt.Println("Withdraw how much : ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > accountBalance {
				fmt.Println("You don't have enough")
				continue
			}
			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greated than 0.")
				continue
			}
			accountBalance -= withdrawAmount
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			fmt.Println("The new balance is", accountBalance)
		default:
			fmt.Println("GoodBye !")
			fmt.Println("Thanks for choosing our bank !")
			return
		}
	}
}
