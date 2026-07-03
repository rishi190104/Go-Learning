package main

import "fmt"

//In structs we define the type of the data
//instead of defining multiple variables we can
//define a single type of struct and use it for
//declaring multiple struct and access them using .

//so the struct is (structure) a custom data type
//that lets you group related data together

//this is how we define the struct
type User struct {
	id     int
	name   string
	number int
}

type Employee struct {
	empId       int
	empName     string
	designation string
}

func updateDesignation(emp *Employee) string {
	emp.designation = "Solution Architect"
	//here we no need to write like (*emp).designation for the pointer as in go struct internally does that
	return emp.designation
}

func main() {
	var user1 = User{
		id:     1,
		name:   "Rishi",
		number: 9082568471,
	}
	//we can also initialize the struct without the field names because go fills fields in order
	//buts its preffered and its also a best practice to have the named fields while initializing
	//because the code is clearer and won't break if the struct definition changes.
	user2 := User{
		2,
		"Rajan",
		999999999,
	}

	fmt.Println(user1)
	fmt.Println(user1.name)

	fmt.Println(user2)
	fmt.Println(user2.number)

	//we can also modify the field of the struct using .
	user2.number = 4567856
	fmt.Println(user2.number)
	fmt.Println(user2)

	//if we dont assign a value go return 0 values
	//so there output for int will be 0 and for string it will empty string
	var user3 User

	fmt.Println(user3)
	//output {0  0}
	//internally its equivalent to User {id: 0, name: "", number:0}

	//structs are value types just like arrays structs are copied when assigned
	user4 := user1
	fmt.Println(user1.name)
	fmt.Println(user4.name)

	//user1 is unchanged because the user4 is the copy of the user1

	//if we want to modify the original we struct we can use pointers

	// type Employee struct {
	// 	empId int
	// 	empName string
	// 	designation string
	// }

	// func updateDesignation(emp Employee) {
	// 	emp.designation = "Solution Architect"
	// }

	employee1 := Employee{
		empId:       909,
		empName:     "Rishi",
		designation: "Technical Lead",
	}

	updateDesignation(&employee1)

	fmt.Println(employee1)

}
