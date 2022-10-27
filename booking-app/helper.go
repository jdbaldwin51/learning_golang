package main

import (
	"fmt"
	"strings"
    "time"
)

// capital first letter of function name "exports" that function to be available in another package -- also can export variables by doing the same
func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
    isValidName := len(firstName) >= 2 && len(lastName) >= 2
    isValidEmail := strings.Contains(email, "@")
    isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
    return isValidName, isValidEmail, isValidTicketNumber
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
    time.Sleep(10 * time.Second)
    var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
    fmt.Println("############")
    fmt.Printf("Sending ticket:\n %v to email address %v\n", ticket, email)
    fmt.Println("############")
    wg.Done()
}
