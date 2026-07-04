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
	//struct does the deferencing part so we dont need to do the (*emp).designation
	return emp.designation
}

type Student struct {
	name string
	age  int
	std  int
}

//This is how we create a method in for a struct by adding a
// receiver type in the function before the name
//by using the below syntax we add the function method to the struct for now we have added struct Student

// func (s Student) getStudentStandard() int {

// }

//there is a convention in go lang when adding a method in the struct we should use the first letter of the struct

//receiver type
func (s *Student) getStudentStandard() int {
	fmt.Println("This is my standard")
	s.std = 20
	return s.std
}

//As struct is used for OOP in go lang so there should be a constructor

//there is no constructor in the go lang we use the function to create a constructor

type Salary struct {
	s_id     int
	s_amount float32
	s_month  string
}

func displaySalary(s_id int, s_amount float32, s_month string) Salary {
	//initials setup goes here
	yourSalary := Salary{
		s_id:     s_id,
		s_amount: s_amount,
		s_month:  s_month,
	}

	return yourSalary
}

//STRUCT EMBEDDING

type Customer struct {
	c_id     int
	c_name   string
	c_number int
	Product
}

type Product struct {
	p_id    int
	p_name  string
	p_desc  string
	p_price float32
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

	var student1 = Student{
		std: 10,
	}

	fmt.Println(student1.getStudentStandard())
	fmt.Println(student1)

	// this work as a constructor
	var salary Salary = displaySalary(1, 22222.22, "july")
	fmt.Println(salary)

	//when we know we will be using the struct only once we can use the below syntax as well

	MySelf := struct {
		name string
		age  int
	}{"Rishi", 22}

	fmt.Println(MySelf)

	//below is the example of the struct embedding

	//we can use declare both the struct differently and then use the with one we have to embed
	//or else we can directly write them inside the struct we need to embed

	var productDetails1 = Product{
		p_id:    1112,
		p_name:  "Dumbell",
		p_desc:  "This is workout equippment",
		p_price: 1000.02,
	}

	var customerDetails1 = Customer{
		c_id:     1,
		c_name:   "Rishi Singh",
		c_number: 9082568471,
		Product:  productDetails1,
	}

	fmt.Println(customerDetails1)

	var customerDetails2 = Customer{
		c_id:     1,
		c_name:   "Clive Alphonso",
		c_number: 9082568471,
		Product: Product{
			p_id:    67,
			p_name:  "Football",
			p_desc:  "This is red football",
			p_price: 300.00,
		},
	}

	fmt.Println(customerDetails2)
	fmt.Println(customerDetails2.Product)
	fmt.Println(customerDetails2.Product.p_name)
}
