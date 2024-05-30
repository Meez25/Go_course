package main

import "fmt"

func main() {
	age := 32 // regular variable

	agePointer := &age

	fmt.Println("Age:", *agePointer)

	MutateAdultYears(&age)

	fmt.Println(age)
}

func MutateAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18
}
