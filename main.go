package main

import (
	"fmt"
	"go-learning/helper"
	"sync"
)

const conferenceTickets int = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50

// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

var wg = sync.WaitGroup{}

// struct is better than map because can mix diffrent type of data
type UserData struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
}

func main() {

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {

		firstName, lastName, email, userTickets := getUserInput()

		//validate input
		isValidName, isValidEmail, isValidRemainingTicket :=
			helper.ValidateInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidRemainingTicket {

			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1) // add how much thread you want to be  waited by main Thread
			go sendTicket(userTickets, firstName, lastName, email)
			// go keyword enable goroutines , and run this line of code on diffrent thread
			// and that apply the concurrency concept

			firstNames := getFirstNames()

			fmt.Printf("The firstNames already booked a tickets:%v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. come back next year.")
				break
			}

		} else {

			if !isValidName {
				fmt.Println("firstname or lastname is too short!")
			}
			if !isValidEmail {
				fmt.Println("email address entered doesn't contain  @ sign")
			}
			if !isValidRemainingTicket {
				fmt.Println("number of ticket you entered is invalid")
			}

		}

	}
	wg.Wait() // wait when thread done
}

func greetUsers(confName string, confTickets int, remainingTicket uint) {
	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T , conferenceName is %T",
	// 	conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcom to %v booking application\n", confName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", confTickets, remainingTicket)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {

	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your userName:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your lastName:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of Tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {

	remainingTickets = remainingTickets - uint(userTickets)

	// var userData = make(map[string]string) //use make to set Map as empty
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10) // here you want to convert from uint to string

	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		userTickets: userTickets,
	}
	bookings = append(bookings, userData)

	fmt.Printf("%v\n", bookings)
	fmt.Printf("Thank you %v %v for booking  %v tickets. You will receive a confirmation email at %v\n",
		firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
