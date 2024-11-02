package main

import (
	"booking-app/shared"
	"fmt"
	"strconv"
)

var conferenceName string = "Go Conference"

const conferenceTickets = 50

var remainingTickets int = 50
var bookings = make([]map[string]string, 0)

//creating empty list of maps
//0 here denotes the initial size of the slice which increases as we increase the element

func main() {

	greetUsers()

	for remainingTickets > 0 {

		firstName, lastName, email, userTickets := userInput()

		isValidEmail, isValidTicketNumber := shared.Validation(email, userTickets, remainingTickets)

		//if user wants to book more tickets than remaining tickets
		if isValidEmail && !isValidTicketNumber {

			// var bookings []string

			bookingTickets(userTickets, firstName, lastName, email)

			firstNames := fname()
			fmt.Printf("These are all our bookings: %v\n", firstNames) //only provides first name of the people who bought tickets

			if remainingTickets == 0 {
				fmt.Println("Our conference is fully booked. Come back next year.")
				break
			}

		} else if userTickets == remainingTickets {
			fmt.Println("Sorry all the tickets have been sold out")
		} else {
			if !isValidEmail {
				fmt.Printf("Invalid email.\n")
			}

			if !isValidTicketNumber {
				fmt.Printf("Number of tickets you entered is invalid.\n")
			}
			fmt.Printf("Your input data is invalid\n\n")
			continue

		}

	}

}

func greetUsers() {
	// var conference = conferenceName
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend the conference\n")

}

func fname() []string { //here the seond []string denotes the firstname is also a slice of string
	firstNames := []string{}
	//this lines below defines whole logic of getting first name of the user
	for _, booking := range bookings { //_ is the blank identifier (ignores the unused variable ) used when you want

		// var names = strings.Fields(booking)                    //splits the booking string into words based on whitespace
		firstNames = append(firstNames, booking["First Name"]) //append the first name to the list

	}
	return firstNames
}

func userInput() (string, string, string, int) {

	var firstName string
	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	var lastName string
	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	var email string
	fmt.Println("Enter your email")
	fmt.Scan(&email)

	var userTickets int
	fmt.Println("Enter number of tickets you want to buy")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookingTickets(userTickets int, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//creating a map for user

	var userData = make(map[string]string)
	userData["First Name"] = firstName
	userData["Last Name"] = lastName
	userData["Email"] = email
	userData["Number of Tickets"] = strconv.FormatInt(int64(userTickets), 10)

	bookings = append(bookings, userData)
	//bookings contain list with all the key value pairs

	fmt.Printf("List of bookings is %v\n", bookings)
	fmt.Printf("Thank You %v %v for booking %v tickets. You will receive the email of your tickets at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remain for the conference %v.\n", remainingTickets, conferenceName)

	return
}
