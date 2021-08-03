package main

import (
	"fmt"
	"log"
)

func main() {
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}

	animals := []string{"dog", "fish", "cat", "mouse", "bird"}

	for i, animal := range animals {
		log.Println(i, animal)
	}

	creatures := make(map[string]string)
	creatures["dog"] = "Snoopy"
	creatures["cat"] = "Tom"

	for creatureType, name := range creatures {
		log.Println(creatureType, name)
	}

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"Mike", "Davis", "mike.davis@abc.co.uk", 40})
	users = append(users, User{"Jan", "Smith", "js@test.com", 49})
	users = append(users, User{"Ben", "Adams", "ben.adams123@ben.co.uk", 20})
	users = append(users, User{"Kenny", "Brown", "kenny.brown@something.co.jp", 35})

	for _, user := range users {
		fmt.Println(user.FirstName)
	}
}
