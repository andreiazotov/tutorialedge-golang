package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Summary:
myChannel := make(chan int) - creates myChannel which is a channel of type int
channel <- value - sends a value to a channel
value := <- channel - receives a value from a channel

Program doesn’t immediately terminate.
This is because the act of sending to and
receiving from a channel are blocking.
main() function blocks until it receives
value from channel.
*/

func calculacteValue(values chan int) {
	value := rand.Intn(10)
	fmt.Println("Calculated value -> ", value)
	values <- value
}

func calculacteValue2(c chan int) {
	value := rand.Intn(77)
	fmt.Println("Calculated value -> ", value)
	time.Sleep(1000 * time.Millisecond)
	c <- value
	fmt.Println("Only Executes after another goroutine performs a receive on the channel")
}

func main() {
	values := make(chan int)
	defer close(values)

	go calculacteValue(values)
	value := <-values
	fmt.Println(value)

	bufferedChannel := make(chan int, 2)
	defer close(bufferedChannel)
	go calculacteValue2(bufferedChannel)
	go calculacteValue2(bufferedChannel)

	result := <-bufferedChannel
	fmt.Println(result)

	time.Sleep(1000 * time.Millisecond)
}
