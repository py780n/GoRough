package main

import "fmt"

type user struct {
	name string
	id   int
}

type programingskill struct {
	user
	language string
	points   int
}

type Notaprogramer struct {
	user
	skillset string
}

func (skill programingskill) asigntask() {
	fmt.Println("Programner name: ", skill.name)
	fmt.Println("Language : ", skill.language)

}

func (skill Notaprogramer) asigntask() {
	fmt.Println("Not a programer")
	fmt.Println("Skill set : ", skill.skillset)

}

type toassign interface {
	asigntask()
}

func usertype(info toassign) {
	info.asigntask()
}

func main() {
	emp1 := new(programingskill)
	emp1.name = "John"
	emp1.id = 99
	emp1.language = "Golang"
	emp1.points = 07

	emp2 := new(Notaprogramer)
	emp2.name = "Tim"
	emp2.id = 99
	emp2.skillset = "Linux"

	//passing Struct type "programingskill" to the interface function
	usertype(emp1)

	//passing Struct type "Notaprogramer" to the interface function
	usertype(emp2)

}
