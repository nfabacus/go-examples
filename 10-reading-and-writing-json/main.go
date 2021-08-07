package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {

	// Converting json string to struct
	myJson := `
	[
		{
			"first_name": "Mike",
			"last_name": "Smith",
			"hair_color": "Brown",
			"has_dog": true
		},
		{
			"first_name": "Mike",
			"last_name": "Smith",
			"hair_color": "Green",
			"has_dog": true
		}
	]`
	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJson), &unmarshalled) // []byte("Some string value") will convert the string to slice of bytes.
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled)

	// converting struct to json
	var mySlice []Person

	var m1 Person
	m1.FirstName = "John"
	m1.LastName = "McDonald"
	m1.HairColor = "red"
	m1.HasDog = false

	var m2 Person
	m2.FirstName = "Json"
	m2.LastName = "Pears"
	m2.HairColor = "black"
	m2.HasDog = false

	mySlice = append(mySlice, m1)
	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "   ") // for api response, not likely to use MarshalIndent
	if err != nil {
		log.Println("error marshalling", err)
	}
	fmt.Println(string(newJson))
}
