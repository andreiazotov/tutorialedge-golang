package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("--------------------")

	for {
		fmt.Print("-> ")

		// If you are only needing single character
		// input then use ReadRune()
		text, _ := reader.ReadString('\n')

		text = strings.Replace(text, "\r\n", "", -1)
		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello")
		}
	}
}
