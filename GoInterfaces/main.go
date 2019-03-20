package main

import "fmt"

// Guitarist interface
type Guitarist interface {
	PlayGuitar()
}

// BaseGuitarist struct
type BaseGuitarist struct {
	Name string
}

// AcousticGuitarist struct
type AcousticGuitarist struct {
	Name string
}

// PlayGuitar base
func (b BaseGuitarist) PlayGuitar() {
	fmt.Printf("%s plays the Bass Guitar\n", b.Name)
}

// PlayGuitar acoustic
func (b AcousticGuitarist) PlayGuitar() {
	fmt.Printf("%s plays the Acoustic Guitar\n", b.Name)
}

// if function expects an empty interface,
// then you can typically pass anything
// into this function
func myFunc(a interface{}) {
	fmt.Println(a)
}

func main() {
	var myAge int
	myAge = 25
	myFunc(myAge)

	var player BaseGuitarist
	player.Name = "Paul"
	player.PlayGuitar()

	var player2 AcousticGuitarist
	player2.Name = "Ringo"
	player2.PlayGuitar()

	var guitarists []Guitarist
	guitarists = append(guitarists, player)
	guitarists = append(guitarists, player2)

	guitarists[0].PlayGuitar()
	guitarists[1].PlayGuitar()
}
