package main

import "fmt"

//now the below function only takes the slice of int as a
//argument but what if we don't know what can be the type of slice
//there we can use generics

//T is a convention for the generics but we can use any letter or name here
func printSlices[T any](s []T) {
	for _, value := range s {
		fmt.Println(value)
	}
}

//if we want the function to accept only the string and int there also we can use generics

func acceptStrorInt[T string | int](arg T) {
	fmt.Println("I only accept data of type string and int", arg)
}

//here we are passing the generics in the struct as s_phone can accept string and int both
type Student[T string | int] struct {
	s_name  string
	s_phone T
}

//instead of any we can also use the interface{} => this acts same as the any

//TO LEARN: need to learn more about this
func typeComparable[T comparable](s T) {

}

func main() {
	nums := []int{1, 2, 3}
	names := []string{"rishi", "clive", "beta"}
	//using generics we can pass any type of data here
	printSlices(names)
	printSlices(nums)

	name := "Rishi"
	age := 22
	// isMale := true

	//so now here if we pass the argument as the type string or int it doesn't complains
	//but as we pass the bool it complains as bool does not satisfy string | int (bool missing in string | int)
	acceptStrorInt(name)
	acceptStrorInt(age)
	// acceptStrorInt(isMale)
}
