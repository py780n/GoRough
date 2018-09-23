package main

import "fmt"

func main() {
	//FOR Loop Method-1
	fmt.Println("### Method 1######")
	index := 1
	for index <= 10 {
		fmt.Println(index)
		index = index + 1
	}
	//FOR  loop Method-2
	fmt.Println("\n### Method 2######")
	for index := 10; index >= 0; index-- {
		fmt.Println(index)
	}
	//FOR  loop Method-3
	fmt.Println("\n### Method 3######")
	for dndex, limit := 100, 90; dndex >= limit; dndex-- {
		fmt.Println(dndex, limit)
	}

}
