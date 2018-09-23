package main

import "fmt"

func message() {
	str := recover()
	fmt.Println(str)

}

func main() {
	defer message()
	panic("this is panic message")

}
