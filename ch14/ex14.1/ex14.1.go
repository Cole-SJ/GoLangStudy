package main

import "fmt"

func main() {
	var a int = 500
	var p *int

	p = &a

	fmt.Printf("%p\n", p)
	fmt.Printf("%d\n", *p)

	*p = 200

	fmt.Printf("%d\n", a)
}
