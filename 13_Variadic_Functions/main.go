package main

import "fmt"

//vardiac functions can take multiple arguments
//the parameters ...int means zero or more integers
// when we write like this numbers ...int we get slice thats why we can iterate over it
func sum(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9))
	//if we want to pass a slice to the variadic function we can that too

	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//this works as spread operator in js
	fmt.Println(sum(nums...))

}
