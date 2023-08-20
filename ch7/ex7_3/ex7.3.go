package main

import "fmt"

func PrintAvgScore(name string, math int, eng int, history int) {
	total := math + eng + history
	avg := total / 3

	fmt.Println("name =", name)
	fmt.Println("avg =", avg)
}

func main() {

	PrintAvgScore("김일등", 80, 74, 95)
	PrintAvgScore("송이등", 88, 92, 53)
	PrintAvgScore("박삼등", 78, 73, 78)

}
