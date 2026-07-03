// In closure the inner function has the access and value of the outer function or parent function
// variable even after the function is executed

//also we can say the inner scope can access the outer scope variable even after there execution

package main

import (
	"fmt"
)

func fullName(fname string) func() string {
	return func() string {
		lname := "Singh"
		return fname + " " + lname
	}
}

func add() func() int {
	num1 := 1000

	return func() int {
		num2 := 17666345
		return num1 + num2
	}
}

func main() {
	name := fullName("Rishi")
	fmt.Println(name())

	addResult := add()
	fmt.Println(addResult())
}
