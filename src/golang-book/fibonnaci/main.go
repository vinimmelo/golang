package main

import (
	"fmt"
)

// Recursive approach to resolve fibonacci calculation
func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	result := fibonacci(10)
	fmt.Println(result)
}
