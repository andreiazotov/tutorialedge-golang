package main

import "fmt"

var name string
var initCounter int
var whatIsThe = answerToLife()

func answerToLife() int {
	return 42
}

func init() {
	fmt.Println("Called First in Order of Declaration")
	initCounter++
}

func init() {
	fmt.Println("Called second in order of declaration")
	initCounter++
}

func init() {
	fmt.Println("This will get called on main initialization")
	name = "MyName"
	whatIsThe = 0
}

/*
In this scenario AnswerToLife() will run
before init() function as WhatIsThe variable
is declared before our init() function is called.
*/
func main() {
	fmt.Println("My Wonderful Go Program")
	fmt.Printf("Name: %s\n", name)
	if whatIsThe == 0 {
		fmt.Println("It's all a lie.")
	}
	fmt.Printf("Init Counter: %d\n", initCounter)
}
