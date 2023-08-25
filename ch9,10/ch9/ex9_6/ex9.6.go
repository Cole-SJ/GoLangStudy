package main

import "fmt"

func getMyAge() (age int, ok bool) {
	return 33, true
}

func main() {

	if age, ok := getMyAge(); ok && age < 20 {
		fmt.Println("you are young", age)
	} else if ok && age < 30 {
		fmt.Println("Nice age", age)
	} else if ok {
		fmt.Println("You are so beautiful", age)
	} else {
		fmt.Println("Error")
	}
	//fmt.Println(age)

}
