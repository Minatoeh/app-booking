package helper

import "strings"

// ValidateUserInput checks if the user inputs are valid.
func ValidateUserInput(userTickets uint, firstName string, lastName string, userMail string, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(userMail, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
