package main

import "fmt"

func main() {
	fmt.Println("Studying Arrays in GoLang")

	var arr [5]int   //declaring a empty array int with 5 spaces   // array needs set the size and type, not can be changed
	fmt.Println(arr) // [0 0 0 0 0]

	arr[0] = 1       // assing a value for index 0
	fmt.Println(arr) // [1 0 0 0 0]

	arrayOfNumbers := [5]int{1, 2, 3, 4, 5} // declaring and initializing an array
	fmt.Println(arrayOfNumbers)             // [1 2 3 4 5]

	arrayPartilized := [5]int{1, 2}     // declaring an array partializing the size
	fmt.Printf("%v\n", arrayPartilized) // [1 2 0 0 0]

	arraySizeAutoCalc := [...]string{"a", "b", "c"} // declaring an array with automatic size calc (need use "" for string)
	fmt.Printf("%v\n", arraySizeAutoCalc)           // [a b c] // size 3
}
