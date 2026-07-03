package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	//range in go are use to iterate over the data types like string, array, slices and maps

	var name string = "Rishi"
	// when we iterate like this on string for the value it returns unicode for the string
	for key, value := range name {
		fmt.Println(key, value)
	}

	//to get the string in the response we need to explicitly convert it using string function
	for key, value := range name {
		fmt.Println(key, string(value))
	}

	var arr = [4]int{10, 20, 30, 40}
	//as the for returns 2 values here like key and the value but wherever we need the value only and the
	// go compiler complain for the key is not used then we can use a underscore _
	for _, val := range arr {
		fmt.Println(val)
	}
	//iterating over slice using range
	var slice = []string{"Rishi", "Rajan", "Sheela", "Ramapati"}

	for key, val := range slice {
		fmt.Println(key, val)
	}

	slices.Reverse(slice)
	fmt.Println(slice)

	//iterating over maps using range

	var mps = map[string]string{
		"name1": "Rishi",
		"name2": "Rajan",
		"name3": "Sheela",
		"name4": "Ramapati",
	}

	for key, value := range mps {
		fmt.Println(key, value)
	}
	// return a arrays of key but not in the order
	//this creates a slice as well
	var mpsKeys = slices.Collect(maps.Keys(mps))
	fmt.Println(mpsKeys)

	// if we want in the order we can use sort

	slices.Sort(mpsKeys)
	fmt.Println(mpsKeys)

}
