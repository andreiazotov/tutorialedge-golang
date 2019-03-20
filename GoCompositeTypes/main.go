package main

import "fmt"

// Person struct
type Person struct {
	name string
	age  int
}

// Team struct
type Team struct {
	name    string
	players [2]Person
}

func main() {
	days := [...]string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}

	fmt.Println(days[0])
	fmt.Println(days[5])

	weekdays := days[0:5]
	fmt.Println(weekdays)

	youtubeSubscribers := map[string]int{
		"AAA":  2222,
		"ABCD": 4444,
		"MMBB": 123321,
	}
	fmt.Println(youtubeSubscribers["AAA"])

	elliot := Person{name: "Elliot", age: 24}
	elliot.age = 18

	var myTeam Team
	fmt.Println(myTeam)

	players := [...]Person{Person{name: "Forest"}, Person{name: "Gordon"}}
	celtic := Team{name: "Celtic FC", players: players}
	fmt.Println(celtic)
}
