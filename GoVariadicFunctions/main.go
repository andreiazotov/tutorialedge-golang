package main

import "fmt"

/*
Variadic Functions are not limited to strings,
we can use any variation of composite or basic types
*/
func variadicFunction(args ...string) {
	fmt.Println(args)
}

func variadicFunctionAny(args ...interface{}) {
	fmt.Println(args...)
}

func main() {
	variadicFunction("xxx", "ddd")
	variadicFunctionAny(1, "aaa", []int{1, 2})
}
