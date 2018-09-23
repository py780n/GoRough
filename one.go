package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func(chn <-chan int) {
		i := <-chn
		fmt.Println(i)
	}(ch)
	go func(chn chan<- int) {
		chn <- 44
	}(ch)
	fmt.Scanln() //"scanln" is used to make the main wait without exiting .
	fmt.Println("DONE")
}
