package main

import (
	"fmt"
	"sync"
)

/*
We first call .Add(1) on our WaitGroup to set
the number of goroutines we want to wait for,
and subsequently, we call .Done() within any
goroutine to signal the end of its’ execution.
Note - You need to ensure that you call
.Add(1) before you execute your goroutine.
*/
func myFunc(wg *sync.WaitGroup) {
	fmt.Println("Enter Goroutine")
	wg.Done()
}

func main() {
	fmt.Println("Enter The Program")

	var wg sync.WaitGroup
	wg.Add(1)
	go myFunc(&wg)
	wg.Wait()

	fmt.Println("Finished Execution")
}
