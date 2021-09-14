package main

import (
	"fmt"
	"log"
)

var x = "Value from outside main!"

func main() {
	fmt.Println("Hello World")
	fmt.Println("x is", x)

	var whatToSay string
	var myNum int

	whatToSay = "Hello, everyone!"
	fmt.Println(whatToSay)

	myNum = 7

	fmt.Println(myNum + 1)

	myValue := "My Param"
	fmt.Println("myValue is ", myValue)
	whatWasSaid, otherThingSaid := saySomething(myValue)
	fmt.Println("myValue is ", myValue, ": no change")

	fmt.Println("The function returned ", whatWasSaid, otherThingSaid)

	greeting := "Greeting with Hello!"
	log.Println(greeting)

	// passing a pointer (reference) of a value
	changeUsingPointer(&greeting)
	log.Println(greeting)
	log.Println("X is now", x)

	fmt.Printf("name is %q, othername is %q, onemorename is %v\n", "Mike", "Bob", "Kate")
	fmt.Printf("Boolean is %v\n", true)
	fmt.Printf("nummer: %2d\n", 31)
	var age = 27
	fmt.Printf("age is %2d\n", age)

}

func saySomething(myParam string) (string, string) {
	fmt.Println(myParam)

	// This will not change myParam outside of this function.
	myParam = myParam + " updated!!"
	// myValue = "Yo!" will not work either, as it does not use lexical scoping, unlike javascript.
	// But, you can use value from outside main func as below.
	fmt.Println("x is", x)
	x = "Changed value"
	fmt.Println("x is", x)
	return myParam, "else"
}

func changeUsingPointer(s *string) {
	log.Println("s is set to ", s)
	// newValue := "I am a new value!"
	*s = "I am a new value!"
}
