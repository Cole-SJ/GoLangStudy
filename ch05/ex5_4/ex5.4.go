package main

import "fmt"

func main() {

	str := "Hello\tGo\t\tWorld\n\"Go\"is Awesome!\n"
	fmt.Print(str)
	fmt.Printf("%s", str)
	fmt.Printf("%q", str) //그대로 출력
}
