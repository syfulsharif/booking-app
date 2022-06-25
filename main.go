package main

import "fmt"

func main() {

	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50 //Using uint will stop the value from going negative

	fmt.Printf("conferenceName is %T, conferenceTickets is %T and remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)
	//%T is a place holder for showing data type in go.

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and currently we have %v tickets available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	//%v is placeholder for showing Value in go.

	var userName string
	var userTickets int
	// If we don't put value in the variable immediately, we need to specify the type
	//ask user for theier name
	userName = "Tom"
	userTickets = 5
	fmt.Printf("User : %v has booked %v tickets.\n", userName, userTickets)

}
