package main

// STRUCTS - which stand for structure, useful for a map with mixed data types
// CONCURRENCY

import (
	"booking-app-8/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50

var bookings = make([]UserData, 0)

var wg = sync.WaitGroup{}

/*
	type keyword is used to create structs ->

comparable to classes in other languages
*/
type UserData struct { // for custom data types
	firstName, lastName, email string
	NumberOfTickets            uint
}

// Greet Users
func greetUsers() {
	fmt.Printf("Welcome to our %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avaliable.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend.")
}

// Get Slice with First Names
func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
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
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		NumberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v Remaining tickets for %v.\n", remainingTickets, conferenceName)
	fmt.Printf("These are all the first names in the bookings%v.\n", getFirstNames())
}

func sendTicket(userTickets uint, firstName, lastName, email string) {
	// Assume that this is a slower process - Ex. generating pdf , sending the attachment - 10 second delay
	time.Sleep(10 * time.Second) // Stops execution of thread for 10 seconds
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##################")
	fmt.Printf("Sending ticket: %v\n to email address %v\n", ticket, email)
	fmt.Println("##################")
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
				wg.Add(1)
				go sendTicket(userTickets, firstName, lastName, email) // go keyword abstracts away the thread creation
				/*
				You would pass 2 to Add if you had another user
				*/
				fmt.Printf("The list of bookings is: %v\n", bookings)
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
