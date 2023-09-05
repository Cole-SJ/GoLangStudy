package main

import "fmt"

func main() {
	math := 80
	eng := 74
	history := 95

	fmt.Println("math + eng + history =", (math+eng+history)/3)

	math = 88
	eng = 92
	history = 53

	fmt.Println("math + eng + history =", (math+eng+history)/3)

	math = 78
	eng = 73
	history = 78

	fmt.Println("math + eng + history =", (math+eng+history)/3)

}
