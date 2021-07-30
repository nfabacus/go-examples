package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	var whatToSay string
	var myNum int

	whatToSay = "Hello, everyone!"
	fmt.Println(whatToSay)

	myNum = 7

	fmt.Println(myNum + 1)

	whatWasSaid, otherThingSaid := saySomething()

	fmt.Println("The function returned ", whatWasSaid, otherThingSaid)
}

func saySomething() (string, string) {
	return "something", "else"
}
