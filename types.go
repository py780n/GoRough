package main

import "fmt"

var thirdname = "Raj" //scope

const constantVariable string = "this is a constant variable"

func main() {
	// &&  operator
	fmt.Println("=AND=")
	fmt.Println("true && true=", true && true)
	fmt.Println("false && false=", false && false)
	fmt.Println("true && false=", true && false)
	fmt.Println("false && true=", false && true)

	//|| operator
	fmt.Println("=OR=")
	fmt.Println(true || true)
	fmt.Println(false || false)
	fmt.Println(true || false)
	fmt.Println(false || true)
	// !
	fmt.Println("=NOT=")
	fmt.Println(!true)
	fmt.Println(!false)

	var name = "First String" // we need not mention the variable type as Go can infer from value assigned
	fmt.Println(len(name))
	name += " Second string" //  append the second string using  '+='
	fmt.Println(name)
	shorthandvariable := "This is a short hand variable" //shorthand
	fmt.Println(shorthandvariable)

	// declaring a constant varable
	const constantVariable1 string = "this is a constant variable inside the function"
	fmt.Println(constantVariable1)
	cont()
}

func cont() {
	fmt.Println(constantVariable)
	var ( //defining multiple variables
		a = "somu"
		b = 99
		c = 22.7
	)
	fmt.Println(a, b, c)

}
