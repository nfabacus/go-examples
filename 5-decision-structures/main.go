package main

import "fmt"

func main() {
	myNum := 1000
	isTrue := true

	if myNum > 99 && isTrue {
		fmt.Println("myNum is greater than 99 and isTrue is set to true")
	} else {
		fmt.Println("other response")
	}

	myVar := "something else"

	switch myVar {
	case "dog":
		fmt.Printf("dog is set to dog")
	case "cat":
		fmt.Printf("dog is set to cat")
	default:
		fmt.Println("This is default")
	}
}
