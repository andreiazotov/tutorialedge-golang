package main

import "fmt"

func myFunc(firstName string, lastName string) (string, error) {
	return firstName + " " + lastName, nil
}

func addOne() func() int {
	var x int
	return func() int {
		x++
		return x + 1
	}
}

func main() {
	fullName, err := myFunc("John", "Doe")
	if err != nil {
		fmt.Println("Error occured")
	}
	fmt.Println(fullName)

	function := addOne()
	fmt.Println(function()) // 2
	fmt.Println(function()) // 3
	fmt.Println(function()) // 4
	fmt.Println(function()) // 5
}
