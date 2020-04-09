package main

import "fmt"

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}

func main() {
	var name, request string
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hss"}
	animals:=map[string]Animal{"cow": cow, "bird":bird, "snake":snake}
	for {
		fmt.Print("Please enter the name and request about the animal: ")
		fmt.Scan(&name)
		fmt.Scan(&request)
		switch request {
		case "eat":
			animals[name].Eat()
		case "move":
			animals[name].Move()
		case "speak":
			animals[name].Speak()
		}
	}

}
