package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("coneference is %T, remaing tickets is %T, conference tickets is %T\n", conferenceName, remainingTickets, conferenceTickets)

	fmt.Println("Get your tickets here! We have", conferenceTickets, "tickets available.")
	fmt.Println("Welcome to the ", conferenceName, "booking application! We have ", remainingTickets, " tickets available.")
	fmt.Printf("Some text with a variable: %v\n", conferenceName)

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		//ask for user name

		//Reading user input
		fmt.Println("Please enter your name: ")
		fmt.Scan(&firstName)

		fmt.Println("Please enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Please enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Please enter the number of tickets you want to buy: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets

		var bookings []string
		//Can also create slice like this: bookings := []string{}

		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v for booking %v tickets. You will recieve a confirmation at your %v email address\n", bookings[0], userTickets, email)
		fmt.Printf("We have , %v, tickets available.\n", remainingTickets)
		fmt.Printf("These are all the bookings: %v\n", bookings)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
	}

}
