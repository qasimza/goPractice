package main

// ARRAYS AND SLICES - Using Slices

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to our %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avaliable.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend.")

	// Slice is an abstraction of an Array, it is also a more flexible & efficient way in many cases
	var bookings []string
	/* ALTERNATIVE WAYS OF DECLARING A SLICE:
	var bookings = []string{}
	bookings := []string{}
	*/
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

	remainingTickets = remainingTickets - userTickets

	//Appending to a slice
	bookings = append(bookings, firstName+" "+lastName)

	/*
		ACCESSING SLICE VALUES IS THE SAME AS
		fmt.Printf("The whole array: %v\n", bookings)
		fmt.Printf("The first value: %v\n", bookings[0])
		fmt.Printf("Array Type: %T\n", bookings)
		fmt.Printf("Array Length (Size): %v\n", len(bookings)) // len is a built-in function to get the size of an array
	*/

	fmt.Printf("These are all the values in the bookings slice %v.\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v Remaining tickets for %v.\n", remainingTickets, conferenceName)
}
