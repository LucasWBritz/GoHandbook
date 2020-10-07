package main

import "fmt"

func main() {
	// declaring variables:
	// name first, then variable type:
	var myVar string

	// multiplea variable declarion
	// To see all available variable types check: https://golang.org/pkg/go/types/#pkg-variables
	var (
		myVar2 string
		myVar3 int
	)

	myVar = "a"
	myVar2 = "b"
	myVar3 = 10

	// When to use var or :=
	// := if you want to create a variable and initialize it.
	myAge := 28

	// use var if you want to initialize a var with his zero value or
	// if you want to declare multiple variables of the same type
	var a, b, c string
	a = "A"
	b = "B"
	c = "C"
	// See more here: https://stackoverflow.com/a/53404567

	// print all variables
	fmt.Println(myVar)
	fmt.Println(myVar2)
	fmt.Println(myVar3)

	fmt.Printf("%v, %v, %v with age %d", a, b, c, myAge)
}
