package main

import (
	"fmt"
	"math"
)

// Struct Circle
type Circle struct {
	x, y, r float64
}

// Method area
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	c := Circle{2, 2, 5}
	fmt.Println(c.area())
}
