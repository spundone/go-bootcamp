package main

import "fmt"

func main () {
	const confName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = conferenceTickets
	fmt.Println("Welcome to our", confName, "booking app")
	fmt.Println("Get your tickets here to attend")

	fmt.Println("ConferenceTickets is set to", conferenceTickets, "and remainingTickets is set to", remainingTickets)

	var userName string
	var userEmail string
	var userTickets int

	fmt.Println("Enter your name:")
	fmt.Scan(&userName)
	fmt.Println("Enter your email:")
	fmt.Scan(&userEmail)
	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	fmt.Printf("User %v booked %v tickets and email is %v\n", userName, userTickets, userEmail)

	remainingTickets = remainingTickets - uint(userTickets)
	fmt.Println("Remaining tickets are", remainingTickets) 
}