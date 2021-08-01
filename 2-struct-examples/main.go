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

func printFirstName(m User) string {
	return m.FirstName
}

// you can assign function to struct easily as below:
func (m *User) printLastName() string {
	return m.LastName
}

func main() {
	var user1 User
	user1.FirstName = "Tony"
	user1.LastName = "Davis"
	user1.PhoneNumber = "123 123 123"

	user2 := User{
		FirstName: "Bob",
		LastName:  "Jones",
	}

	fmt.Println(user1.FirstName, user1.LastName)
	fmt.Println(user1.PhoneNumber)

	fmt.Println(user2.FirstName, user2.LastName, user2.BirthDate)
	fmt.Println(user2.PhoneNumber)

	fmt.Println(printFirstName(user1))
	fmt.Println(user1.printLastName())
	fmt.Println(user2.printLastName())

}
