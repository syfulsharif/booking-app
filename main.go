package main

import "fmt"

func main() {

	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50 //Using uint will stop the value from going negative

	var firstName string
	var lastName string
	var email string
	var ticketsToBuy uint

	// fmt.Printf("conferenceName is %T, conferenceTickets is %T and remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)
	// //%T is a place holder for showing data type in go.

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and currently we have %v tickets available\n", conferenceTickets, remainingTickets)

	var bookings []string
	//bookings = append(bookings, firstName+" "+lastName)
	for {
		fmt.Println("Enter Your first name: ")
		fmt.Scan(&firstName) // &"variableName" is pointer, it addresses the memory location of variable saved

		fmt.Println("Enter Your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter Your email address: ")
		fmt.Scan(&email)

		fmt.Println("How many tickets would you like to buy?")
		fmt.Scan(&ticketsToBuy)

		remainingTickets -= ticketsToBuy
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for buying %v tickets, a confirmation mail will be sent at %v\n", firstName, lastName, ticketsToBuy, email)
		fmt.Printf("%v tickets remained for %v\n", remainingTickets, conferenceName)

	}

	// //fmt.Printf("The whole slice: %v\n", bookings)
	// fmt.Printf("The first Valule: %v\n", bookings[0])
	// fmt.Printf("Slice Type: %T\n", bookings)
	// fmt.Printf("The length of the slice %v\n", len(bookings))
}
