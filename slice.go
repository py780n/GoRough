package main

import "fmt"

func main() {
	// Declaring a slice
	slice1 := []string{"I", "AM", "GO", "PROGRAMER"}
	//slice1[0] = "Go"
	fmt.Println("Length of array : ", len(slice1))
	fmt.Println("Capacity of array : ", cap(slice1))
	fmt.Println("Array : ", slice1)

	//Declaring  Slice using built-in make() function
	slice2 := make([]int, 10, 15) // creates a slice with length of '10' and capacity of '15'
	for index := 0; index < len(slice2); index++ {
		slice2[index] = index
	}
	fmt.Println("Length of array : ", len(slice2))
	fmt.Println("Capacity of array : ", cap(slice2))
	fmt.Println("Array : ", slice2)
}
