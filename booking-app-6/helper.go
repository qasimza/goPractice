package main

import "strings"

// Validate User Input
func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

/*
In order to run main now, all files should be specified as such -
go run main.go helper.go
Or specify a folder location
go run . (. specifies current folder)
*/
