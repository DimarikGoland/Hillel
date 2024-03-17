package main

import (
	"flag"
	"fmt"
	"sort"
)

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
	distance := 1000
	fmt.Printf("Distance: %d\n ", distance)
	fish, turtle, tiger := CreateTeam()
	fishSpeed := flag.Int(fish.name, 1, "Speed F")
	turtleSpeed := flag.Int(turtle.name, 1, "Speed T")
	tigerSpeed := flag.Int(tiger.name, 1, "Speed Tig")

	flag.Parse()

	fmt.Printf("Speed of fish: %v\n", *fishSpeed)
	fmt.Printf("Speed of turtle: %v\n", *turtleSpeed)
	fmt.Printf("Speed of tiger: %v\n", *tigerSpeed)
	_, runSpeedF := Race{
		Distance: distance,
		fish: Fish{
			name:      fish.name,
			fins:      fish.fins,
			colorSkin: fish.colorSkin,
		},
	}.Start(*fishSpeed)

	animalFish := Animal{
		runSpeed: runSpeedF,
		name:     fish.name,
		voice:    "",
	}
	_, runSpeedT := Race{
		Distance: distance,
		turtle: Turtle{
			name:      turtle.name,
			legs:      turtle.legs,
			colorSkin: turtle.colorSkin,
		},
	}.Start(*turtleSpeed)
	animalTurtle := Animal{
		runSpeed: runSpeedT,
		name:     turtle.name,
		voice:    "",
	}
	_, runSpeedTiger := Race{
		Distance: distance,
		tiger: Tiger{
			name:      tiger.name,
			legs:      tiger.legs,
			colorSkin: tiger.colorSkin,
		},
	}.Start(*tigerSpeed)
	animalTiger := Animal{
		runSpeed: runSpeedTiger,
		name:     tiger.name,
		voice:    "",
	}
	var animals []Animal
	animals = append(animals, animalFish, animalTurtle, animalTiger)
	sort.Sort(ByRunSpeed(animals))

	fmt.Printf("%v is win, time off iteration: %d\n ", animals[0].name, animals[0].runSpeed)
	fmt.Println(Animal{}.WinnerSay())
	fmt.Printf("Second plays is %v, time off iteration: %d\n ", animals[1].name, animals[1].runSpeed)
	fmt.Printf("Third plays is %v, time off iteration: %d\n ", animals[2].name, animals[2].runSpeed)
	fmt.Printf(Animal{}.LoserSay())
}

func (a Animal) WinnerSay() string {
	a.voice = "I am best of the best off the best! "
	return a.voice
}
func (a Animal) LoserSay() string {
	a.voice = "I am win next time! "
	return a.voice
}
func CreateTeam() (Fish, Turtle, Tiger) {
	fish := Fish{
		"Nemo",
		8,
		"grey",
	}
	turtle := Turtle{
		"Tortilla",
		4,
		"black",
	}
	tiger := Tiger{
		"Balu",
		4,
		"Yellow-black",
	}
	return fish, turtle, tiger
}

func (r Race) Start(animal int) (race Race, iter int) {
	var i int

	for {
		if r.Distance <= 0 {
			break
		}
		r.Distance = r.Distance - animal
		i++

	}
	return race, i
}

type ByRunSpeed []Animal

func (b ByRunSpeed) Len() int           { return len(b) }
func (b ByRunSpeed) Less(i, j int) bool { return b[i].runSpeed < b[j].runSpeed }
func (b ByRunSpeed) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
