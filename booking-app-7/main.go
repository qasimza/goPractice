package main

// FUNCTIONS AND PACKAGE LEVEL VARS

import (
	"booking-app-7/helper" //You have to explicitly state the path for a non built in function because go will look for the directory
	"fmt"
	"strings"
)

// The following are available all throughout the package, ie. Package Level Arrays

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50

var bookings []string

// Great Users
func greetUsers() {
	fmt.Printf("Welcome to our %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avaliable.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend.")
}

// Get Slice with First Names
func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

// Get User Input
func getUserInput() (string, string, string, uint) {
	var firstName, lastName, email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Enter the number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

// Book Tickets
func bookTicket(userTickets uint, firstName, lastName, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v Remaining tickets for %v.\n", remainingTickets, conferenceName)
	fmt.Printf("These are all the first names in the bookings%v.\n", getFirstNames())
}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked up. Come back next year.")
				break
			} else {
				bookTicket(userTickets, firstName, lastName, email)
			}

		} else {
			if !isValidName {
				fmt.Println("ERROR: First and Last Name should be at least 2 characters.")
			}
			if !isValidEmail {
				fmt.Println("ERROR: Your email is invalid. It does not contain @")
			}
			if !isValidTicketNumber {
				fmt.Println("ERROR: INvalid ticket number")
			}
			//fmt.Printf("Your input data is invalid.")
			continue // break the current iteration instead of the entire loop
		}

	}
}
