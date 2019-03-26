package main

import (
	"fmt"
	"time"
)

func compute(input int) {
	for i := 0; i < input; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("goroutine tutorial")

	go compute(10)
	go compute(10)

	// anonymous function concurrent using `go`
	go func() {
		fmt.Println("Executing my Concurrent anonymouse function")
	}()

	var s string
	fmt.Scanln(&s)
}
