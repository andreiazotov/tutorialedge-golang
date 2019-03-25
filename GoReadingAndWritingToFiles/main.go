package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	mydata := []byte("Data to write to a file\n")
	if err := ioutil.WriteFile("localfile.data", mydata, 0777); err != nil {
		fmt.Println(err)
	}

	data, err := ioutil.ReadFile("localfile.data")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	f, err := os.OpenFile("localfile.data", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString("additional data write to the file\n"); err != nil {
		panic(err)
	}

	data, err = ioutil.ReadFile("localfile.data")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}
