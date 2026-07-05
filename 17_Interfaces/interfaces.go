//there is no implement keyword in GO as it is in java or C++

//If a type has a required method, it automatically satisfies the interface

//like below we have a interface name paymenter which has a method pay
//and below them we have the struct razorpay and stripe with the pay method so
//this automatically satisfies the paymenter interface
//This is called as Implicit Implementation

//interface removes duplicate code
//without interface we would need a different methods for all of the payment methods
//now we only have the generic function makePayment() and for all the provider we just
//need to create the pay method that we have defined in the payementer interface

//interface enables polymorphism > one function works with many types

//Easier testing

//suppose we need to send email
/*
type SendEmailer interface {
sendEmail(to, message string) string
}

for actual email
type Email struct {}

func (e Email) sendEmail(to, message string) string {

}

for test email
type TestEmail struct {}

func (t TestEmail) sendEmail(to, message string) string {

}

type Send struct {
emailSend: SendEmailer
}

func (s Send) SendingEmail(to, message string) {
s.emailSend.sendEmail(to, message)
}
*/

package main

import "fmt"

type paymenter interface {
	pay(amount int)
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount int) {
	// fmt.Println("Making Payment of", amount)
	// razorPayGw := razorpay{}
	// razorPayGw.pay(amount)

	// stripeGw := stripe{}
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount int) {
	fmt.Println("Making payment using razorpay", amount)
}

type stripe struct{}

func (s stripe) pay(amount int) {
	fmt.Println("Making payent using stripe", amount)
}

type SendEmailer interface {
	sendEmail(to, message string)
}

// for actual email
type Email struct{}

func (e Email) sendEmail(to, message string) {
	fmt.Println("Sending Actual Email To" + " " + to + " " + message)
}

// for test email
type TestEmail struct{}

func (t TestEmail) sendEmail(to, message string) {
	fmt.Println("Sending Test Email To" + " " + to + " " + message)
}

type Send struct {
	emailSend SendEmailer
}

func (s Send) SendingEmail(to, message string) {
	s.emailSend.sendEmail(to, message)
}

func main() {

	// razorpayGw := razorpay{}
	stripeGw := stripe{}

	doPayment := payment{
		gateway: stripeGw,
	}

	doPayment.makePayment(100)

	testEmail := TestEmail{}
	actualEmail := Email{}

	sendTestEmail := Send{
		emailSend: testEmail,
	}

	sendActualEmail := Send{
		emailSend: actualEmail,
	}

	sendTestEmail.SendingEmail("Rishi Singh", "I hope you are doing well")
	sendActualEmail.SendingEmail("Rishi Singh", "I hope you are doing well")

}
