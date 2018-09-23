package main

import (
	"fmt"
	"sort"
)

//Returning single value function
func TopScore(dataList []float64) float64 {
	sort.Float64s(dataList) //sort in-built function is used to sort the list
	return dataList[len(dataList)-1]
}

//Returning Multiple Values funtion
func average(dataList []float64) (float64, int) {
	//This function calculates the average for the given list of numbers
	total := 0.0
	for _, value := range dataList {
		total += value
	}
	return total / float64(len(dataList)), len(dataList) //Returning Multiple Values
}

//Variadic Functions
func totalMarks(nums ...float64) int {
	total := 0.0
	for _, value := range nums {
		total += value
	}
	return int(total)
}

//recursion function
func factorial(value int) int {
	if value == 0 {
		return 1
	}
	return value * factorial(value-1)
}

//function created to use defer
func endthepro() {
	fmt.Println("the Program ends here")
}

// panic recover function
func message() {
	str := recover()
	fmt.Println(str)
}

func main() {
	defer endthepro()
	defer message()

	markList := []float64{98, 64, 75, 80}            //slice to hold the marks of the student
	studentAverage, numberofsub := average(markList) // passing the mark list to the average and number of subjects .
	studentTopscore := TopScore(markList)            // passing the mark list to get the top score
	totalscore := totalMarks(markList...)
	fmt.Println("Top Socer :", studentTopscore)
	fmt.Println("Student average :", studentAverage)
	fmt.Println("Number of subjects :", numberofsub)
	fmt.Println("Total Score :", totalscore)

	//Closure :
	X := 100
	incremant := func() int {
		X++
		return X
	}
	increamentby10 := func(val int) int {
		return val + 10
	}
	fmt.Println(increamentby10(X))
	fmt.Println(incremant())

	fmt.Println("Factorial :", factorial(10))
	panic("this is panic message")

}
