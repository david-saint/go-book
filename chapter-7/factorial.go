package main

import "fmt"

func main() {
	fmt.Println("Enter the number you want to find the factorial of: ")
	var number int
	fmt.Scanf("%d", &number)
	fmt.Printf("The factorial of %d is %d", number, factorial(number))
}

func factorial(v int) int {
	if v == 0 {
		return 1
	}

	return v * factorial(v-1)
}
