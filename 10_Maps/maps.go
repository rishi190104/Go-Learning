//make
//delete
//clear

package main

import "fmt"

func main() {
	m := map[string]string{"name": "rishi"}

	fmt.Println(m["name"])

	m1 := make(map[string]string, 2)

	fmt.Println(m1)
	fmt.Println(len(m1))

	m1["name"] = "Rishi"
	fmt.Println(m1)
	// we can use this to check if the key exists or not
	//if the m1["name"] this key doesn't exist then it return false and for the
	//name it returns zero values if string then "" empty, if int then 0 if bool then false like this
	name, exists := m1["name"]

	fmt.Println(name, exists)

	// iterating over maps

	var marks = map[string]int{
		"student1": 10,
		"student2": 10,
		"student3": 10,
		"student4": 10,
		"student5": 10,
	}
	//output order may vary
	//map order is not guaranteed in Go.
	for key, value := range marks {
		fmt.Println(key, value*2)
	}
	//for deleting a value from a map we use delete function it takes name of the map and key as the arguments
	delete(m, "name")

	fmt.Println(m)
}
