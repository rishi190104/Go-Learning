package main

import "fmt"

func main() {
	//once declared cannot be updated
	const designation string = "Full Stack Engineer"

	//multiple vars declaration at once
	var (
		firstName string = "Rishi"
		lastName string = "Singh"
	)

	//multiple const declaration at once
	const (
		age int = 21
		isWorking bool = true
	)

	fmt.Println(designation)
	fmt.Println(firstName, lastName)
	fmt.Println(age, isWorking)
}