package main

import "fmt"

const PIG int = 0
const COW int = 1
const CHICKEN int = 2

func PrintAnimal(animal int) {
	if animal == PIG {
		fmt.Println("꿀꿀")
	} else if animal == COW {
		fmt.Println("음매")
	} else if animal == CHICKEN {
		fmt.Println("꼬끼오")
	} else {
		fmt.Println("...")
	}
}

func main() {
	PrintAnimal(PIG)
	PrintAnimal(COW)
	PrintAnimal(CHICKEN)
}
