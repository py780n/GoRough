package main

import "fmt"

func main() {
	fmt.Println("\nif/else statemet output\n")
	//example for if/else statement
	/*here we are checking if the number is odd or even .
	Go works in top down method the loop is increamnted as the condition is
	satisfied */
	for value, limit := 1, 25; value <= limit; value++ {
		if value%2 == 0 {
			fmt.Println(value, " is even")
		} else {
			fmt.Println(value, " is odd")
		}
	}

	//example for  if/else if statement
	/*here we are checking if the numner is divisible by 10,5,7 in the same order
	and we print the number in stdout .
	Go works in top down method the loop is increamnted as the condition is
	satisfied */
	fmt.Println("\nif/else if statemet output\n")
	for value, limit := 0, 100; value <= limit; value++ {
		if value%10 == 0 {
			fmt.Println(value, " is div by 10")
		} else if value%5 == 0 {
			fmt.Println(value, " is div by 5")
		} else if value%7 == 0 {
			fmt.Println(value, " is div by 7")
		}
	}
	for val := 0; val <= 5; val++ {
		switch val {
		case 0:
			fmt.Println("zero")
		case 1:
			fmt.Println("one")
		case 4:
			fmt.Println("four")
		default:
			fmt.Println("Not defined")

		}
	}
}
