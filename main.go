package main

import "fmt"

func main() {

	const name string = "André"
	const age int = 10
	const weight float64 = 55.5
	var isCool = true

	isCool = false

	isBeautiful := true
	isBeautiful = false

	fmt.Println("My name is: ", name)
	fmt.Println("My age: ", age)
	fmt.Println("My weight: ", weight)
	fmt.Println("Am I cool? ", isCool)
	fmt.Println("Am I beautiful? ", isBeautiful)
}
