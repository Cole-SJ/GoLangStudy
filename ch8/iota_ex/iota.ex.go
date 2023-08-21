package main

import "fmt"

const (
	RED   int = iota
	BLUE  int = iota
	GREEN int = iota
)

const (
	C1 uint = iota + 1
	C2
	C3
)

const (
	BIG_FLAG1 uint = 1 << iota
	BIG_FLAG2
	BIG_FLAG3
	BIG_FLAG4
)

func main() {
	fmt.Println(RED, BLUE, GREEN)
	fmt.Println(C1, C2, C3)
	fmt.Println(BIG_FLAG1, BIG_FLAG2, BIG_FLAG3, BIG_FLAG4)
}
