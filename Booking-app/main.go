package main

import (
	"fmt"
	"booking-app/helper"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets int = 50
var bookings = make([]map[string]string, 0)

func main() {
	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber :=
			helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are available\n",
		conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}

func getUserInput() (string, string, string, int) {
	var firstName, lastName, email string
	var userTickets int

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets you want to book: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets int, firstName, lastName, email string) {

	remainingTickets -= userTickets

	// create a map for user
	user := make(map[string]string)
	user["firstName"] = firstName
	user["lastName"] = lastName
	user["email"] = email
	user["userTickets"] = fmt.Sprint(userTickets)

	bookings = append(bookings, user)

	fmt.Printf(
		"Thank you %v %v for booking %v tickets. A confirmation email will be sent to %v\n",
		firstName, lastName, userTickets, email,
	)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
