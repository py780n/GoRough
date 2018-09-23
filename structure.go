package main

import (
	"fmt"
	"math"
)

type userinfo struct {
	userName string
	Userid   int
	userMail string
}

type cylinder struct {
	hight  float64
	radius float64
}

//"Embeded strucuture"
type CodingSkill struct {
	userinfo
	cancode bool
	lang    string
}

//method
func (cylindertype *cylinder) area() float64 {
	return (2 * math.Pi * cylindertype.radius * cylindertype.hight) + (2 * math.Pi * cylindertype.radius * cylindertype.radius)

}

type X int

func (y X) dd() {
	fmt.Println(y)
}

func addMailid(detail *CodingSkill) {
	detail.userMail = detail.userName + "@golang.com"
}

func main() {

	//fmt.Println(x)
	empdetail := &userinfo{}
	empdetail.userName = "Josh"
	empdetail.Userid = 99
	empdetail.userMail = "Josh@golang.com"

	fmt.Println(empdetail.userName)
	//fmt.Println

	cyl1 := cylinder{10, 5}
	fmt.Println("Area of cylinder :", cyl1.area()) //  used the "Method" type funciton

	//using embeded type struct
	programer := new(CodingSkill)
	programer.userName = "john"
	programer.Userid = 100
	addMailid(programer)
	programer.cancode = true
	programer.lang = "GO"

	//z := new(X)
	X(10).dd()

	fmt.Println(*programer)

}
