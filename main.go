package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings [] string{}
	//Can also create slice like this: bookings := []string{}
	greetUser(conferenceName)


	fmt.Println("coneference is %v, remaing tickets is %v, conference tickets is %v", conferenceName, remainingTickets, conferenceTickets)

	fmt.Println("Get your tickets here! We have", conferenceTickets, "tickets available.")
	fmt.Println("Welcome to the ", conferenceName, "booking application! We have ", remainingTickets, " tickets available.")
	fmt.Printf("Some text with a variable: %v\n", conferenceName)

	for remainingTickets > 0| len(bookings) < 50 {
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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2  
		isValidEmail := strings.Contains(email, "@") 
		isValidTickets := userTickets > 0 && userTickets <= remainingTickets
		// isInvalidCity := city != "New York" || city != "London"

		//check if there are enough tickets
		if isValidEmail && isValidName && isValidTickets{
			remainingTickets = remainingTickets - userTickets

			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v for booking %v tickets. You will recieve a confirmation at your %v email address\n", bookings[0], userTickets, email)
			fmt.Printf("We have , %v, tickets available.\n", remainingTickets)
			fmt.Printf("These are all the bookings: %v\n", bookings)


			firstNames := getFirstNames(bookings)

			var noTickets bool = remainingTickets == 0
			if noTickets {
				fmt.Println("Sorry, we are sold out!")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("Sorry, your name is invalid, try again.")
			}
			if !isValidEmail {
				fmt.Println("Sorry, your email is invalid, try again.")
			}
			if !isValidTickets {
				fmt.Println("Sorry, the number of tickets you entered is invalid, try again.")
			}
 		}

	}

	city := "London"
	switch city {
	case "New York":
		//execute code for booking in New York
	case "San Francisco":
		//execute code for booking in San Francisco
	case "London", "Paris":
		//execute code for booking in London
	default:
		//execute code for booking in other cities
		fmt.Printf("Sorry, we don't have any events in this\n", city)
	}
}


func greetUser(conferenceName string) {
	fmt.Println("Welcome to the %v conference", conferenceName)
}

func getFirstNames(bookings []string) []string{
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}


