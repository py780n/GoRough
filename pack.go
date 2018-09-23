package main

import (
	"fmt"
	"golang/math"
)

func main() {
	val := []float64{10, 10, 10}
	val2 := math.TopScore(val)
	fmt.Println(val2)

}
