package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

type user struct {
	name string
	id   int
}

func main() {
	cn := make(chan int, 10)
	//us := make(chan user)

	wg.Add(3)
	go func(ch <-chan int) {
		for ele := range ch {
			//i := <-ele
			fmt.Println(ele)
		}
		wg.Done()
	}(cn)

	go func(ch <-chan int) {
		for ele := range ch {
			//i := <-ele
			fmt.Println("new :", ele)
		}
		wg.Done()
	}(cn)

	go func(ch chan<- int) {
		for i := 10; i <= 100; i = i + 10 {
			ch <- i
		}
		close(ch)
		wg.Done()
	}(cn)
	wg.Wait()

}
