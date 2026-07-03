package main

import (
	"errors"
	"fmt"
)

//basic functions

func printName(name string) {
	fmt.Println("Hello", name)
}

// function with some return value
func add(num1 int, num2 int) int {
	return num1 + num2
}

// we can also write a shorter version of the above function if the parameters type is same
// then no need to repeat it we can just write the type at the last
// this is called shorter parameter syntax
func concateName(fname, lname string) string {
	return fname + lname
}

//in go function can return multiple values as well
//so in return type as we mention the types it should be same while returning and calling a function

func arithmaticOperations(num1, num2 int) (int, int, int, int) {
	add := num1 + num2
	sub := num1 - num2
	mul := num1 * num2
	div := num1 / num2

	return add, sub, mul, div
}

//returning error along with response

func errorRes(age int) (int, error) {

	return age, errors.New("This is error")
}

//named return values in func

func namedReturn(length, width int) (area int) {
	area = length * width
	return
}

func main() {

	printName("Rishi")
	add(1, 2)
	//we can also store this in a variable
	addResult := add(19, 1)
	fmt.Println(addResult)
	fmt.Println(concateName("Rishi", "Singh"))

	//for getting the multiple values from a function we need multiple variable
	//the function returns  in the same order as in the function declaration
	addition, substraction, multiplication, division := arithmaticOperations(2, 2)
	fmt.Println(addition, substraction, multiplication, division)
	//if we don't want any return value from the function we can use _ there as in the for loop
	addition1, _, multiplication1, division1 := arithmaticOperations(5, 2)
	fmt.Println(addition1, multiplication1, division1)

	fmt.Println(errorRes(3))

	fmt.Println(namedReturn(3, 3))

	//anonymous functions

	greet := func() {
		fmt.Println("Hello")
	}
	greet()

	//iife

	func() {
		fmt.Println("this is iife")
	}()
}
