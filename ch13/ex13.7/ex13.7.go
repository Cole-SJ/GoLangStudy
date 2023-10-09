package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	A int8
	B int
	C int8
	D int
	E int8
}

func main() {
	user := User{
		A: 1,
		B: 2,
		C: 3,
		D: 4,
		E: 5,
	}

	fmt.Println("unsafe.Sizeof(user) =", unsafe.Sizeof(user))
}
