package main

import "fmt"

func getMyAge() int {
	return 22
}

func main() {
	switch age := getMyAge(); age {
	case 10:
		fmt.Println("TeenAge")
	case 33:
		fmt.Println("Pair 3")
	default:
		fmt.Println("My age is ", age)
	}

}