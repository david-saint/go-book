package main

import "fmt"

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	smallest := -1

	for _, value := range x {
		if smallest == -1 {
			smallest = value
		} else if value < smallest {
			smallest = value
		}
	}

	fmt.Printf("The smallest number is %d", smallest)
}
