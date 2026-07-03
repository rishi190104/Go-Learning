package main

import "fmt"

func main() {
	//in go array has a fixed length
	// we can use it when we know the size of the array
	var nums = [2]int{}
	nums[0] = 1
	nums[1] = 2
	fmt.Println(nums)

	//when we create a array without any values in it so it
	// takes zero values as the value of the array
	// so for string it takes empty strings "", for int it takes 0
	//for bool it takes false

	var names = [3]string{}
	fmt.Println(names, "string")

	var num = [3]int{}
	fmt.Println(num, "int")

	var boolean = [3]bool{}
	fmt.Println(boolean, "bool")

	var floats = [3]float32{}
	fmt.Println(floats, "float32")

	// in go we can create a 2D (2 dimension) arrays as well

	nos := [2][3]int{{3, 3, 3}, {3, 5, 6}}
	fmt.Println(nos)
	//above is 2d that is matrix need to learn more about this

	// it can used when we have a fixed length due to which there is memory optimization
}
