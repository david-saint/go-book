package main

import "fmt"

var name string = "Hello World"

func main() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}

func f() {
	fmt.Println(name)
}
