package main

import (
	"fmt"
	"strings"
)
func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = conferenceTickets
	bookings:= []string{}

	fmt.Printf("Welcome to %v booking application\n",conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are available\n",conferenceTickets,remainingTickets)
	println("Get your tickets here to attend")

	for{
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

		isValidName := len(firstName) >=2 && len(lastName) >=2
		isValidEmail := strings.Contains(email,"@")
		isValidTicketNumber := userTickets > 0 && userTickets <= uint(remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber{
			remainingTickets = remainingTickets - int(userTickets)
			bookings = append(bookings,firstName + " " + lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email address at %v\n",firstName,lastName,userTickets,email)
			fmt.Printf("%v tickets are remaining for %v\n",remainingTickets,conferenceName)

			firstNames := []string{}
			for _,booking := range bookings{
				var names = strings.Fields(booking)
				firstNames = append(firstNames,names[0])
			}
			fmt.Printf("The first name of bookings are : %v\n",firstNames)
			
			if(remainingTickets == 0){
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		}else {
			if !isValidName {
				fmt.Printf("First name or last name you entered is too short\n")
			}
			if !isValidEmail {
				fmt.Printf("Email address you entered doesn't contain @ sign\n")
			}
			if !isValidTicketNumber {
				fmt.Printf("Number of tickets you entered is invalid\n")
			}
		}		
	}
}