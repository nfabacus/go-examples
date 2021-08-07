package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Cat struct {
	Name  string
	Color string
	Age   int
}

func main() {
	dog := Dog{
		Name:  "Bob",
		Breed: "Chihuahua",
	}

	cat := Cat{
		Name:  "Pen",
		Color: "White",
	}

	PrintInfo(&dog)

	PrintInfo(&cat)

}

func PrintInfo(a Animal) {
	fmt.Println("this animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}

// better to pass pointers
func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (d *Cat) Says() string {
	return "Meao!"
}

func (d *Cat) NumberOfLegs() int {
	return 4
}
