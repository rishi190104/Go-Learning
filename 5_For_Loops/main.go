package main

import "fmt"

func main() {
	for i := 0; i <= 3; i++ {
		fmt.Println(i)
	}
	//there is no while loop in the go --- so we write while loop like this
	var i = 0
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	// infinite loop in go
	for {
		fmt.Printf("Infinite loop")
	}
}
