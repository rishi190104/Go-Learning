package main

import "fmt"

func main() {

	//for package level variable use var

	// string
	var name string = "Rishi"
	//shorthand ro declare a variable
	fullName := "Rishi Singh"
	//number
	var age int = 21
	//boolean
	var isWorking bool = true
	//float
	var salary float32 = 3.19

	fmt.Println(name, age, isWorking, salary, fullName)

	var greet string = "HI"
	// left side of shorthand property should be a new variable not a existing variable
	// shorthand variable declaration can only be used inside a function
	// greet := "Hello"
	fmt.Println(greet)
}