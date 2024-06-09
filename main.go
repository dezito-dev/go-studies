package main

import "fmt"

/* ============== Aggregate type ===============*/
/*
 aggregate type is a type of data that groups together other values or variables. These values can be of the same type or different types. In Go, arrays, structs, and slices are examples of aggregate types.
*/

func main() {
	fmt.Println("Studying Arrays in GoLang")

	/* ============== Arrays ===============*/

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

	/* ============== Structs ===============*/
	// structs are a composite data type to group together different data types

	type Person struct {
		name    string
		age     int
		weight  float64
		hobbies []string // it's a slice
	}

	var p1 Person // declaring a variable of type Person

	p1.name = "Andre" // assing values to the struct fields
	p1.age = 21
	p1.weight = 70.5
	p1.hobbies = []string{"Play games", "Read books"}

	fmt.Println(p1) // {Andre 21 70.5 [Play games Read books]}

	// declaring and initializing a struct
	p2 := Person{name: "John", age: 30, weight: 80.5, hobbies: []string{"Play games", "Read books"}}
	fmt.Println(p2) // {John 30 80.5 [Play games Read books]}

}
