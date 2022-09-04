package main

import "fmt"

func main() {

	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("%v tickets remaining out of %v tickets\n", remainingTickets, conferenceTickets)
	fmt.Println("get your ticket to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Print("how many tickets you wanna get: ")
	fmt.Scan(&userTickets)

	fmt.Printf("Thnak you %v %v for puchasing %v tickets for %v \n", firstName, lastName, userTickets, conferenceName)
	fmt.Printf("A confirmation email has been sent to %v \n", email)
}
