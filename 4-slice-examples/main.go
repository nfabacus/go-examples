package main

import (
	"fmt"
	"sort"
)

func main() {
	var mySlice []string
	var myNumbers []int
	otherNumbers := []int{10, 2, 3, 4, 5, 6}
	myStrList := []string{"a", "b", "c", "test"}

	mySlice = append(mySlice, "Ben")
	mySlice = append(mySlice, "Tony")
	mySlice = append(mySlice, "Mary")

	myNumbers = append(myNumbers, 2)
	myNumbers = append(myNumbers, 4)
	myNumbers = append(myNumbers, 1)
	myNumbers = append(myNumbers, 3)

	fmt.Println(mySlice)
	fmt.Println(myNumbers)
	sort.Ints(myNumbers)
	fmt.Println(myNumbers)
	fmt.Println(otherNumbers)
	fmt.Println(otherNumbers[0:3])
	fmt.Println(myStrList)
}
