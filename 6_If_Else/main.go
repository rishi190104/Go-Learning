package main

import "fmt"

func main() {
	age := 1

	if age >= 18 {
		fmt.Println("Adult")
	} else if age >= 13 && age <= 18 {
		fmt.Println("Teen age")
	} else {
		fmt.Println("Kid")
	}

	// decalaring variable in the if statement
	if num := 10; num%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}
}
