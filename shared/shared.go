package shared

import "strings"

func Validation(email string, userTickets int, remainingTickets int) (bool, bool) {
	isValidEmail := strings.Contains(email, "@") //strings package has Contains which checks whether the given value is present or not
	isValidTicketNumber := userTickets < 0 && userTickets <= remainingTickets
	return isValidEmail, isValidTicketNumber
}
