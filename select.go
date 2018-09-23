package main

import (
	"fmt"
	"sync"
)

var wtg = sync.WaitGroup{}

var donCh = make(chan struct{})

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	wtg.Add(2)
	go func(ch1 chan<- int) {
		for i := 0; i <= 10; i++ {
			ch1 <- i
		}
		wtg.Done()
	}(ch1)

	go func(ch2 chan<- int) {
		for j := 11; j <= 20; j++ {
			ch2 <- j
		}
		wtg.Done()
	}(ch2)
	//wtg.Wait()

	go findRange(ch1, ch2)
	donCh <- struct{}{}
	wtg.Wait()
}

//function using select statement
func findRange(chl1 <-chan int, chl2 <-chan int) {
	for {
		select {
		case entry1 := <-chl1:
			fmt.Println("The number between 0 to 10 :", entry1)
		case entry2 := <-chl2:
			fmt.Println("Then number between 11 to 20 :", entry2)
		case <-donCh:
			break
		}
	}

}
