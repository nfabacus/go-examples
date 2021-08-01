package main

import (
	"fmt"
	"time"
)

// struct is a bit like type declaration for an object in typescript.
type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName: "Bob",
		LastName:  "Jones",
	}

	fmt.Println(user.FirstName, user.LastName, user.BirthDate)
	fmt.Println(user.PhoneNumber)
}
