package main

import (
	"github.com/nfabacus/myprogram/helpers"
	"log"
)

func main() {
	log.Println("Hello")

	var myVar = helpers.SomeType{
		TypeName:   "Some name",
		TypeNumber: 7,
	}
	log.Println(myVar.TypeName, myVar.TypeNumber)
}
