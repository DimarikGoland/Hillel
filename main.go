package main

import "fmt"

type Animal struct {
	runSpeed int
	name     string
	voice    string
}
type Turtle struct {
	name      string
	legs      int
	colorSkin string
}
type Tiger struct {
	name      string
	legs      int
	colorSkin string
}

type Fish struct {
	name      string
	fins      int
	colorSkin string
}

type Race struct {
	Distance int
	turtle   Turtle
	tiger    Tiger
	fish     Fish
}

func main() {
	fish, turtle, tiger := CreateTeam()
	fmt.Println(fish, tiger, turtle)
}
func (Animal) WinnerSay() {

}
func (Animal) LoserSay() {

}
func CreateTeam() (Fish, Turtle, Tiger) {
	fish := Fish{
		"nemo",
		8,
		"grey",
	}
	turtle := Turtle{
		name:      "Tortilla",
		legs:      4,
		colorSkin: "black",
	}
	tiger := Tiger{
		name:      "Balu",
		legs:      4,
		colorSkin: "Yellow-black",
	}
	return fish, turtle, tiger
}
