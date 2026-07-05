//GO lang doesn't have any built-in enums as java or C#

//GO uses typed constants

package main

import "fmt"

//like below we create enums for string

type MessageStatus string

const (
	Seen      MessageStatus = "seen"
	Sent                    = "sent"
	Delivered               = "delivered"
	Received                = "received"
)

type ClimateType int

const (
	Cold ClimateType = iota //iota is a untyped in integer, it is zero index and keeps on autoincrement
	//so here value of Cold is 0 and this keeps on increasing for all the variables the Summer value is 1,
	//Monsoon value is 2, Spring value is 3
	Summer
	Monsoon
	Spring
)

func getMessageStatus(status MessageStatus) {
	fmt.Println("Message status is", status)
}

func getClimateStatus(status ClimateType) {
	fmt.Println("Today's weather is", status)
}

func (c ClimateType) String() string {
	switch c {
	case Cold:
		return "cold"
	case Summer:
		return "summer"
	case Monsoon:
		return "monsoon"
	case Spring:
		return "spring"
	default:
		return "unknown"
	}
}

//Go already has an interface
//The fmt package defines an interface like this:

/*
type Stringer interface {
	String() string
}
*/

func main() {
	getMessageStatus(Received)
	//with the string method this will return the string instead of the int
	getClimateStatus(Monsoon)
}
