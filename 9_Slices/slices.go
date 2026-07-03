//make
//copy
//append
//len
//cap
//slice (:)
//slice package
//2d slices

package main

import (
	"fmt"
	"slices"
)

func main() {

	//slices are dynamic they can be expanded there is no fixed length

	nums := []int{} //this returns an empty slice
	fmt.Println(nums)

	//we can also create a slice using make function
	//make function takes 3 arguments
	//syntax make(type, length, capacity)

	names := make([]string, 2, 5)
	fmt.Println(names)

	//to get the length we can use len function
	fmt.Println(len(names))

	//to get the capacity we use cap function
	fmt.Println(cap(names))
	//as soon as we reach the capacity of the slice or we cross
	//the capacity it gets multiplied by 2 automatically

	//to add a value in the slice we use append function
	// it takes 2 args 1st the name of the slice 2nd the value to be added or appended.
	fmt.Println(append(names, "rishi"))

	//as array we can also create a 2d slice as well

	twoDslice := [][]int{{10, 10}, {20, 20}}
	fmt.Println(twoDslice)

	copyTwoDslice := [][]int{}
	//to copy a slice into another variable we use copy function
	//syntax copy(destination, source) => destination in which variable the slice should be copied and source which slice should be copied
	copy(copyTwoDslice, twoDslice)

	//we can also use slice method to get the particular values from start to end
	//syntax [startindex: last index] last is not included

	var counts = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(counts[1:3])
	fmt.Println(counts[:3]) //if we don't mention the start it takes from very start the last index mentioned
	fmt.Println(counts[1:]) //if we don't mention the last index it takes from start mentioned till last

	//slices packages from go also gives us many methods for slice

	var num1 = []int{1, 2, 3}
	var num2 = []int{1, 2, 3}
	fmt.Println(slices.Equal(num1, num2)) // if matches returns true else false

}

//we will be majorly using slice in go server building
