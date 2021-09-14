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

	var someFunction = func() int {
		return 1
	}

	if result := someFunction(); result > 0 { // you can pre-assign variable within if scope (only available in the scope).
		fmt.Printf(" result %d > 0\n", result)
	} else {
		fmt.Printf(" result %d <= 0\n", result)
	}

	myVar := "dog"

	switch myVar {
	case "dog":
		fmt.Printf("myVar is set to dog")
	case "cat":
		fmt.Printf("myVar is set to cat")
	default:
		fmt.Println("This is default")
	}
}
