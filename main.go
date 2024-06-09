package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	// using explicit declaration for variable
	var name string = "André" // string type
	fmt.Printf("Type: %T, Value: %v\n", name, name)

	var age int = 25 // int type
	fmt.Printf("Type: %T, Value: %v\n", age, age)

	var isValid bool = false // boolean type
	fmt.Printf("Type: %T, Value: %v\n", isValid, isValid)

	var weight float64 = 80.1 // float64 type
	fmt.Printf("Type: %T, Value: %v\n", weight, weight)

	// using implicit declaration for variable
	var city = "Floarianopolis" // string type
	fmt.Printf("Type: %T, Value: %v\n", city, city)

	year := 2024 // string type
	fmt.Printf("Type: %T, Value: %v\n", year, year)
}
