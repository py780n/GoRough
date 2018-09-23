package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func add(num1, num2, result *int) {
	*result = *num1 + *num2
	wg.Done()
}

func mult(num1, num2, result *int) {
	*result = *num1 * *num2
	wg.Done()
}

func main() {
	val1 := 10
	val2 := 20
	result1 := 0
	result2 := 0
	wg.Add(2)
	go add(&val1, &val2, &result1)

	go mult(&val1, &val2, &result2)
	wg.Wait()
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result1 * result2)

}
