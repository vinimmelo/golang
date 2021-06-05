package main

import (
	"fmt"
)

// A function that halves a number and return the number
// should return true if it was even or false if it was odd too
func halfANumber(x int) (int, bool) {
	halved := x / 2
	if x%2 == 0 {
		return halved, true
	} else {
		return halved, false
	}
}

func main() {
	fmt.Println(halfANumber(3))
	fmt.Println(halfANumber(4))
}
