package main

import (
	"github.com/nfabacus/myprogram/helpers"
	"log"
)

// in terminal in the root of your project, you do `go mod init github.com/your-github-username/package-name`
// It will create `go.mod` file.
// Then, create a folder e.g. helpers. Then, create a file e.g. helpers.go.  Go will create a file with package helpers

func main() {
	log.Println("Hello")

	var myVar = helpers.SomeType{
		TypeName:   "Some name",
		TypeNumber: 7,
	}
	log.Println(myVar.TypeName, myVar.TypeNumber)
}
