package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	// websites := map[string]string{
	// 	"Google":              "https://google.com",
	// 	"Amazon Web Services": "https://aws.com",
	// }
	// fmt.Println(websites)
	// fmt.Println(websites["Google"])
	//
	// websites["LinkedIn"] = "https://linkedin.com"
	// fmt.Println(websites)
	//
	// delete(websites, "LinkedIn")
	//
	// fmt.Println(websites)

	userNames := make([]string, 2, 5)

	userNames[0] = "Julie"

	userNames[1] = "1"

	userNames = append(userNames, "Yann")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	courses := make(floatMap, 3)

	courses["go"] = 4.7
	courses["react"] = 4.8
	courses["angular"] = 4.7
	// Reallocating memory
	courses["node"] = 4.7

	courses.output()

	for index, value := range userNames {
		fmt.Println(value, index)
	}

}
