package main

import "fmt"

func main() {
	fmt.Println(3*4^7<<2+3*5 == 7)
	fmt.Println((((3 * 4) ^ (7 << 2)) + (3 * 5)) == 7)
	fmt.Println(((12 ^ 28) + 15) == 7)
	fmt.Println((16 + 15) == 7)
	fmt.Println(31 == 7)
}
