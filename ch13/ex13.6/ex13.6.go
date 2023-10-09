package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Age   int32
	Score float64
}

func main() {
	user := User{
		Age:   23,
		Score: 77.2,
	}

	fmt.Println(unsafe.Sizeof(user))
}
