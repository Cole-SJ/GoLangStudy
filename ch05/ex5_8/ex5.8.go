package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	stdin := bufio.NewReader(os.Stdin)

	var a int
	var b int
	n, err := fmt.Scanln(&a, &b)

	if err != nil {
		fmt.Println(err)
		/*readString, err := stdin.ReadString('\n')
		if err != nil {
			print(readString)
			return
		}*/ // 지우면 버퍼에 그대로 스트림이 남아있어 문제가 발생한다.
	} else {
		fmt.Println(n, a, b)
	}
	n, err = fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
		readString, err := stdin.ReadString('\n')
		if err != nil {
			print(readString)
			return
		}
	} else {
		fmt.Println(n, a, b)
	}

}
