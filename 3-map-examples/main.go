package main

import "fmt"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	myMap := make(map[string]string)

	myMap["greeting"] = "Hello!"

	fmt.Println(myMap["greeting"])

	users := make(map[string]User)

	users["person1"] = User{
		FirstName: "John",
	}

	fmt.Println(users["person1"].FirstName)
}

// map is not sorted. You always have to look up by key.
