package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("안녕하세요")
		var number int

		_, err := fmt.Scanln(&number)

		if err != nil {
			fmt.Println("숫자를 입력하세요")

			stdin.ReadString('\n')
			continue
		}

		fmt.Printf("입력하신 숫자는 %d 입니다.\n", number)

		if number%2 == 0 {
			break
		}
	}

	fmt.Println("for문이 종료되었습니다.")
}
