package math

import "sort"

func TopScore(dataList []float64) float64 {
	sort.Float64s(dataList) //sort in-built function is used to sort the list
	return dataList[len(dataList)-1]
}
