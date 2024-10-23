package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets = 50

	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")

	for {

		var firstName string
		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)
		// fmt.Printf("First Name: %v\n", firstName)

		var lastName string
		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)
		// fmt.Printf("Last Name: %v\n", lastName)

		var email string
		fmt.Println("Enter your email")
		fmt.Scan(&email)
		// fmt.Printf("Email: %v\n", email)

		var userTickets int

		fmt.Println("Enter number of tickets you want to buy")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets

		var bookings []string
		//bookings[0] = firstName + " " + lastName
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank You %v %v for booking %v tickets. You will receive the email of your tickets at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remain for the conference %v.\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("These are all our bookings: %v\n", bookings)
	}
}
