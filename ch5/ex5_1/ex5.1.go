package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	var f float64 = 32799438743.8297

	fmt.Print("a:", a, "b:", b)
	fmt.Println("a:", a, "b:", b, "f:", f) //지숫값으로 표현된다.. 실수값의 기본 서식은 %f가 아닌 %g이기 때문이다.
	fmt.Printf("a: %d b: %d f: %f\n", a, b, f)

}
