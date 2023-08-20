package main

import "fmt"

func Divide(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return
	}
	result = a / b
	success = true
	return
}

func main() {
	c, success := Divide(9, 3)
	fmt.Println("c, success =", c, success)
	fmt.Println("&success =", &success)
	d, success := Divide(9, 0)
	fmt.Println("d, success =", d, success)
	fmt.Println("&success =", &success)

	var a = 3
	f, a := 2, 4
	fmt.Println("a,f=", a, f)
}
