package main // declare the main package
/*
main = standard package name for where the go application is written
Run the following command to create module for the application -
	go mod init [APP_NAME]
go.mod should be in the parent dir, same level as main.go
*/

import "fmt"

/*
need to explicitly import each package, HOVER over for link to documentation
MUST USE IMPORTED PACKAGES
*/

//main is declaration of entry point of application, only 1 main per app (1 entry point)
func main() {
	var conferenceName = "Go Conference"
	//  SPECIFIC TO GOLANG - On creating/defining a variable, it must be used, otherwise an error is thrown
	const conferenceTickets = 50
	// Use const keyword to prevent the value of a var from changing
	var remainingTickets uint = 50
	// SHORTHAND SYNTAX - remainingTickets := 50            This cannot be used for const

	/*
		You can also define the type explicitly (not necessary) in above statement
		as such var remainingTickets uint = 50
	*/

	/* PRINTING TYPE OF A VARAIBLE
	fmt.Printf("Types:\n conferenceName is %T\n conferenceTickets is %T\n remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)
	*/

	// Use Printf for formatted strings

	//fmt.Println("Welcome to our", conferenceName, "booking application")
	fmt.Printf("Welcome to our %v booking application.\n", conferenceName)

	//fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Printf("We have total of %v tickets and %v are still avaliable.\n", conferenceTickets, remainingTickets)

	fmt.Println("Get your ticket here to attend.")

	/* Go is a statically typed language, you need to tell the Go compiler the type when declaring the variable
	   If you assign a value, it automatically decides the type
	*/

	var firstName, lastName, email string
	var userTickets uint

	// Ask User for information
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Enter the number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v Remaining tickets for %v.\n", remainingTickets, conferenceName)
}
