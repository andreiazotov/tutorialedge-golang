package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// Users struct
type Users struct {
	XMLName xml.Name `xml:"users"`
	Users   []User   `xml:"user"`
}

// User struct
type User struct {
	XMLName xml.Name `xml:"user"`
	Type    string   `xml:"type,attr"`
	Name    string   `xml:"name"`
	Social  Social   `xml:"social"`
}

// Social struct
type Social struct {
	XMLName  xml.Name `xml:"social"`
	Facebook string   `xml:"facebook"`
	Twitter  string   `xml:"twitter"`
	Youtube  string   `xml:"youtube"`
}

func main() {
	xmlFile, err := os.Open("data.xml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Open file successfully")
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)
	var users Users
	xml.Unmarshal(byteValue, &users)

	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Type: " + users.Users[i].Type)
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	}
}
