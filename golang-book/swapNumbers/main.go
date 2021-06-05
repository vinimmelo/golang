package main

import (
	"fmt"
)

// Simple swap function to training pointers
func swap(x *int, y *int) {
	z := *y
	*y = *x
	*x = z
}

func main() {
	x := 12
	y := 25
	swap(&x, &y)
	fmt.Println(x, y)
}
