package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex
var balance int

func init() {
	balance = 1000
}

func withdraw(value int, done chan bool) {
	mutex.Lock()
	fmt.Printf("Withdrawing %d from account with balance: %d\n", value, balance)
	balance -= value
	mutex.Unlock()
	done <- true
}

func deposit(value int, done chan bool) {
	mutex.Lock()
	fmt.Printf("Depositing %d to account with balance: %d\n", value, balance)
	balance += value
	mutex.Unlock()
	done <- true
}

func main() {
	fmt.Println("go mutex")

	done := make(chan bool)
	go withdraw(700, done)
	go deposit(500, done)
	<-done
	<-done

	fmt.Printf("New balance: %d\n", balance)
}
