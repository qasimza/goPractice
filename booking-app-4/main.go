package main

// FOR LOOP, IF ELSE, SWITCH

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to our %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avaliable.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend.")

	var bookings []string

	for {
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

		// USER INPUT VALIDATION

		// First and Last Names are at least 2 characters long
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		// Email must contain @
		isValidEmail := strings.Contains(email, "@")
		// tickets must be greater than 0 and less than or equal to remaining tickets
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets

			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v Remaining tickets for %v.\n", remainingTickets, conferenceName)

			firstNames := []string{}

			// FOLLOWING IS THE FOR REACH CONTSRTUCT IN GO, range gives index, value pairs for each item in a slice/array
			for _, booking := range bookings { // _ replaces index -> BLANK IDENTIFIER FOR UNUSED VARIABLES TO BE IGNORED
				var names = strings.Fields(booking) //Fields is a built-in function that splits a string at space
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These are all the first names in the bookings%v.\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked up. Come back next year.")
				break // Stops execution of loop

				/*
					You could also define the condition at the beginning of the for loop instead of this construct
					for remainingTickets > 0 {
						// code
					}
				*/
			}

			/*
				ALTERNATIVE SYNTAX
				Store flag for no tickets remaining as under
				noTicketsRemaining bool:= remainingTickets == 0
				if remainingTickets{
				}
			*/

			// Use else if for checking other conditions

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

		city := "London"
		switch city {
		case "New York":
			// execute code for booking  NY conference
		case "Berlin":
			// execute code for booking Berlin coference
		case "Singapore", "Hon Kong":
			// consolidate multiple values and execute code for both through here
		default:
			// execute default case - INVALID DATA - no  match found
		}

	}
}
