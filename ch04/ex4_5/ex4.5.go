package main

import "fmt"

func main() {
	var a int16 = 3456
	var c int8 = int8(a)

	fmt.Println(a)
	fmt.Println(c) //16개의 비트 중 앞에 있는 8개의 비트가 사라지기 때문에 -128이 된다.
}
