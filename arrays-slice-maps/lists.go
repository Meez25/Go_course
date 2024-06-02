package main

import "fmt"

func main() {
	fmt.Println("TASK 1")
	myHobbies := [3]string{"Programming", "D&D", "Reading"}
	fmt.Println(myHobbies)

	fmt.Println("---------------")

	fmt.Println("TASK 2")
	fmt.Println(myHobbies[0])
	fmt.Println(myHobbies[1:3])

	fmt.Println("---------------")

	fmt.Println("TASK 3")
	slice := myHobbies[:2]
	fmt.Println(slice)
	slice = myHobbies[0:2]
	fmt.Println(slice)

	fmt.Println("---------------")

	fmt.Println("TASK 4")
	subSlice := slice[1:3]
	fmt.Println(subSlice)

	fmt.Println("---------------")

	fmt.Println("TASK 5")
	courseGoals := []string{"Get better in Go", "Have a better understanding of Computer science"}
	fmt.Println(courseGoals)

	fmt.Println("---------------")

	fmt.Println("TASK 6")
	courseGoals[1] = "Computer Science !"
	courseGoals = append(courseGoals, "Third Goal !")
	fmt.Println(courseGoals)

	fmt.Println("---------------")

	fmt.Println("TASK 7")
	productList := []Product{NewProduct("Titre 1", 1, 10.99), NewProduct("Titre 2", 2, 6.99)}
	productList = append(productList, NewProduct("Titre 3", 3, 1.99))
	fmt.Println(productList)

}

type Product struct {
	title string
	id    int
	price float64
}

func NewProduct(title string, id int, price float64) Product {
	return Product{
		title: title,
		id:    id,
		price: price,
	}
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
