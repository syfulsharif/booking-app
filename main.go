package main

import "fmt"

func main() {

	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50 //Using uint will stop the value from going negative

	var firstName string
	var lastName string
	var email string
	var ticketsToBuy int

	// fmt.Printf("conferenceName is %T, conferenceTickets is %T and remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)
	// //%T is a place holder for showing data type in go.

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and currently we have %v tickets available\n", conferenceTickets, remainingTickets)

	fmt.Println("Enter Your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter Your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter Your email address: ")
	fmt.Scan(&email)

	fmt.Println("How many tickets would you like to buy?")
	fmt.Scan(&ticketsToBuy)

	//%v is placeholder for showing Value in go.

	// var userName string
	// var userTickets int
	// // If we don't put value in the variable immediately, we need to specify the type
	// //ask user for theier name
	// userName = "Tom"
	// userTickets = 5
	// fmt.Printf("User : %v has booked %v tickets.\n", userName, userTickets)

	fmt.Printf("Thank you %v %v for buying %v tickets, a confirmation mail will be sent at %v\n", firstName, lastName, ticketsToBuy, email)

}
