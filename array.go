package main

import "fmt"

func main() {

	var arr [5]int // declatring a array of  length 5
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4

	fmt.Println(arr)

	arr2 := [3]int{6, 7, 8} //shorthand declaration of array methos - 1
	fmt.Println(arr2)

	arr3 := [...]int{9, 10, 11} //shorthand without mentioning the lenght ,method-2
	fmt.Println(arr3)
}
