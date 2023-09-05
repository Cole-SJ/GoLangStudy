package main

import (
	"fmt"
	"math"
)

func equal(a float64, b float64) bool {
	return math.Nextafter(a, b) == b
}

func main() {
	var (
		a float64 = 0.1
		b float64 = 0.2
		c float64 = 0.3
	)

	fmt.Printf("%0.18f + %0.18f = %0.18f\n", a, b, a+b)
	fmt.Printf("%0.18f == %0.18f : %v\n", c, a+b, equal(a+b, c))

	a = 0.0000000000004
	b = 0.0000000000002
	c = 0.0000000000007

	fmt.Printf("%g == %g : %v\n", c, a+b, equal(a+b, c))
}
