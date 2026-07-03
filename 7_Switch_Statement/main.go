package main

import "fmt"

func main() {
	day := "monday"

	switch day {
	case "monday":
		fmt.Println("It's monday")
	case "tuesday":
		fmt.Println("It's tuesday")
	default:
		fmt.Println("It's sunday")
	}

	//switch without an expression
	switch {
	case day == "monday":
		fmt.Println("It's monday")
	case day == "tuesday":
		fmt.Println("It's tuesday")
	default:
		fmt.Println("It's sunday")
	}

	//switch with multiple value in one case
	var days string = "sunday"
	switch days {
	case "sunday", "saturday":
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's weekday")
	}

	//fallthrough

	//Unlike C/C++, Go does not automatically fall through to the next case.

	//even if the case matches and we still want to move to the next case we can use fallthrough there.

	switch 1 {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two")
	}

	//in go switch doesn't needs a break keyword to stop it does it automatically behind the scenes
}
