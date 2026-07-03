package main

import "fmt"

func nums(number int) {
	number = 10
}

func nums1(number *int) {
	*number = 10
}

func main() {
	fmt.Println("This is pointers.....")

	var name string = "Rishi"
	//This prints the value of the variable
	fmt.Println(name)
	//This prints the address where the variable is store
	//so to get the address where the variable is store we
	//use & before the variable name
	fmt.Println(&name)

	//In the below scenario the value of x should get updated after the function is executed
	//so in the function the value should be get updated from 9 to 10 but it didn't happend
	// This happens because Go always passes arguments by value.
	// When nums(x) is called, Go creates a copy of x and passes that copy to the function.
	var x int = 9
	nums(x)
	fmt.Println(x)

	//we solve the above problem using go here so what happens is
	//here we pass the address of the x1 variable by using &
	//the function receives and there we have declare the pointer
	//like number *int (means a pointer of int)
	// Inside the function, *number dereferences the pointer.
	// This lets us access the original value stored at that memory address.
	// Assigning *number = 10 updates the original variable.
	// We use pointers when a function needs to modify the original variable
	// or when we want to avoid copying large structs.

	// Go still passes arguments by value.
	// However, in this case it copies the pointer (memory address),
	// not the integer itself.
	// Both the caller and the function point to the same variable,
	// so modifying *number changes the original value.

	var x1 int = 9
	nums1(&x1)
	fmt.Println(x1)

	// &    ==> Get the memory address of a variable
	// *    ==> Dereference a pointer (read or write the value at that address)
	// *int ==> A pointer to an int (used when declaring a pointer)

	//go always copy arguments but with pointers it copies the address and without the pointer
	//it copies the value

}
