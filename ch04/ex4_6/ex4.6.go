package main

import "fmt"

var g int = 10

func main() {

	var m int = 20
	{
		var s int = 50
		fmt.Println(m, s, g)
	}

	//m = s + 20 변수의 scope를 벗어나서 때문에 에러가 발생한다.
}
