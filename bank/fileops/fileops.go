package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(filename string) (float64, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return 1000, errors.New("Failed to find file.")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse value")
	}

	return balance, nil
}

func WriteFloatToFile(value float64, filename string) {
	valueText := fmt.Sprint(value)
	err := os.WriteFile(filename, []byte(valueText), 0644)
	if err != nil {
		fmt.Println("There was an issue with the save of the value", err.Error())
		return
	}
	fmt.Println("File saved")
}
