package main

import "fmt"

func main() {
	// x := [5]float64{98, 93, 77, 82, 83}

	// slices
	// var x []float64
	// x := make([]float64, 5)

	// var total float64 = 0
	// for _, value := range x {
	// 	total += value
	// }
	// fmt.Println(total / float64(len(x)))

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
}
