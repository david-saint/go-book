package main

import "fmt"

func main() {
	fmt.Println("Enter the number you want to find the fibonacci sequence for: ")
	var number int
	fmt.Scanf("%d", &number)
	fmt.Printf("The fibonacci of %d is %d", number, fibonacci(number))
}

func fibonacci(x int) (r int) {
	if x == 0 {
		r = 0
	} else if x == 1 {
		r = 1
	} else {
		r = fibonacci(x-1) + fibonacci(x-2)
	}
	return
}
