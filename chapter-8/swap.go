package main

import "fmt"

func main() {
	x := 1
	y := 2
	swap(&x, &y)
	fmt.Println(x, y)
}

func swap(xp *int, yp *int) {
	*xp, *yp = *yp, *xp
}
