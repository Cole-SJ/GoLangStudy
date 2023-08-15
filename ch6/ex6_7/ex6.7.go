package main

import "fmt"

const epsilon = 0.0000001

func equal(a float64, b float64) bool {
	if a > b {
		if a-b <= epsilon {
			return true
		} else {
			return false
		}
	} else {
		if b-a <= epsilon {
			return true
		} else {
			return false
		}
	}
}

func main() {

	var (
		a float64 = 0.1
		b float64 = 0.2
		c float64 = 0.3
	)

	fmt.Printf("%0.18f + %0.18f = %0.18f\n", a, b, a+b)
	fmt.Printf("%0.18f == %0.18f : %v\n", c, a+b, equal(a+b, c))

}
