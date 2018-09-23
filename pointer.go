package main

import "fmt"

//function that handles pointer
func makeValuezero(val *int) {
	*val = 0 //"*" here is used to dereference to assigne new value
}
func make999(val *int) {
	*val = 999
}

func sawp(val1 *int, val2 *int) {
	temp := *val1
	*val1 = *val2
	*val2 = temp
}

func main() {
	currentValue := 99
	fmt.Println("Initial  Value: ", currentValue)
	fmt.Println("address of the Initial value :", &currentValue)
	makeValuezero(&currentValue) //"&currentValue" is the adderess that is passed to the funtion
	fmt.Println("New Value: ", currentValue)
	fmt.Println("address of the New value :", &currentValue)

	usingNewfunc := new(int)
	make999(usingNewfunc)
	fmt.Println(*usingNewfunc)

	num1 := 90
	num2 := 89
	fmt.Println(num1, num2)
	sawp(&num1, &num2)
	fmt.Println(num1, num2)

}
