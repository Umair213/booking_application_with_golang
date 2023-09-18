package main

import (
	"fmt"
	"strings"
)

var bookings = make([]userData, 0)

type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	greetUser(conferenceName, conferenceTickets, remainingTickets)

	for {
		var firstName string
		var lastName string
		firstName, lastName = nameInformationAndValidation(firstName, lastName)

		var email string
		email = emailInformationAndValidation(email)

		var userTickets uint
		userTickets = userTicketsAndValidation(userTickets, remainingTickets)

		remainingTickets = ticketsBooking(remainingTickets, userTickets, firstName, lastName, email)

		fmt.Printf("%v %v booked %v tickets for the %v. %v's Email is %v\n", firstName, lastName, userTickets, conferenceName, firstName, email)
		fmt.Printf("%v tickets are remaining\n", remainingTickets)

		printFirstNames()

		if remainingTickets == 0 {
			fmt.Println("We are all booked!")
			break
		}
	}
}

func greetUser(conference string, totalTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to the %v\n", conference)
	fmt.Printf("We have total of %v, tickets and, %v tickets are remaining\n", totalTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func nameInformationAndValidation(firstName string, lastName string) (string, string) {
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	for len(firstName) < 2 || len(lastName) < 2 {
		fmt.Println("Your name is not valid. Please re-enter")
		fmt.Print("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)
	}
	return firstName, lastName
}

func emailInformationAndValidation(email string) string {
	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	for !strings.Contains(email, "@") {
		fmt.Println("Your email is not valid. Please re-enter")
		fmt.Print("Enter your email again: ")
		fmt.Scan(&email)
	}
	return email
}

func userTicketsAndValidation(userTickets uint, remainingTickets uint) uint {
	fmt.Print("Enter number of tickets you want: ")
	fmt.Scan(&userTickets)

	for userTickets > remainingTickets {
		fmt.Printf("You cannot exceed the limit of %v\n", remainingTickets)
		fmt.Println("Enter number of tickets you want again: ")
		fmt.Scan(&userTickets)
	}
	return userTickets
}

func ticketsBooking(remainingTickets uint, userTickets uint, firstName string, lastName string, email string) uint {
	// create a map for user
	var userData = userData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	return remainingTickets
}

func printFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}
