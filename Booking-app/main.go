package main

import "fmt"
func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = conferenceTickets
	var bookings [50]string

	fmt.Printf("Welcome to %v booking application\n",conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are available\n",conferenceTickets,remainingTickets)
	println("Get your tickets here to attend")

	
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("Enter your first name:- ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name:- ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address:- ")
	fmt.Scan(&email)

	fmt.Print("Enter your number of tickets you want to book:- ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - int(userTickets)
	bookings[0]=firstName + " " + lastName

	fmt.Printf("The whole array: %v\n",bookings)
	fmt.Printf("The first value: %v\n",bookings[0])
	fmt.Printf("The type of bookings is %T\n",bookings)
	fmt.Printf("The length of bookings is %v\n",len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email address at %v\n",firstName,lastName,userTickets,email)
	fmt.Printf("%v tickets are remaining for %v",remainingTickets,conferenceName)
}