package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wtgp = sync.WaitGroup{}

func main() {
	/*var msg = "hello"
	go func(msg string) {
		fmt.Println(msg)
	}(msg)
	msg = "yellow"
	time.Sleep(100 * time.Millisecond)*/
	runtime.GOMAXPROCS(4)
	wtgp.Add(2)
	go hello()
	//wtgp.Wait()
	fmt.Println("Last line")
	//wtgp.Add(2)
	go yellow()
	wtgp.Wait()

}

func hello() {
	fmt.Println("hello")
	wtgp.Done()
}

func yellow() {
	fmt.Println("number of thread; ", runtime.GOMAXPROCS(-1))
	wtgp.Done()
}
