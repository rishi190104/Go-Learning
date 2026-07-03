package main

import "fmt"

func main(){
	// var name string = "Rishi Singh"

	firstName := "Rishi"
	lastName := "Singh"
	// like this we can decalre multiple const variables altogether
	const (
		name = "Rishi"
		age = 21
		isWorking = true
	)

	fmt.Println(age)

	fmt.Println(firstName)
	fmt.Println(lastName)
}