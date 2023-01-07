package helper

import "strings"

// export a function by making the first letter capital
func ValidateInput(firstName string, lastName string, email string,
	userTickets uint, remainingTickets uint) (bool, bool, bool) {

	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidRemainingTicket := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidRemainingTicket

}
