package main

import (
    "fmt"
    "sync"
)

var conferenceName = "Go Conference"
const conferenceTickets = 50
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
    firstName string
    lastName string
    email string
    numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

    greetUsers(conferenceName, conferenceTickets, remainingTickets)

// remainingTickets > 0 && len(bookings) < 50

    firstName, lastName, email, userTickets := getUserInput()
    isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

    if  isValidName && isValidEmail && isValidTicketNumber {

        bookTicket(userTickets, firstName, lastName, email)

        wg.Add(1)
        go sendTicket(userTickets, firstName, lastName, email)

            firstNames := getFirstNames()
            fmt.Printf("The first names of bookings are: %v\n", firstNames)

        if remainingTickets == 0 {
              // end program
              fmt.Println("Our conference is booked out, try next year!")
          }

    } else {
        if !isValidName {
            fmt.Println("first name or last name entered is too short")
        }
        if !isValidEmail {
            fmt.Println("email address is invalid")
        }
        if !isValidTicketNumber {
            fmt.Println("number of tickets is invalid")
        }
    }
    wg.Wait()
}

func greetUsers(conferenceName string, conferenceTickets int, remainingTickets uint) {
    fmt.Printf("Welcome to %v booking application\n", conferenceName)
    fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
    fmt.Println("Get your tickets here!")
}

func getFirstNames() []string {
    firstNames := []string{}
    for _, booking := range bookings { // index needs be identified but since it is unused, it can be an underscore
        firstNames = append(firstNames, booking.firstName)
    }
    return firstNames
}

func getUserInput() (string, string, string, uint) {
    var firstName string
    var lastName string
    var email string
    var userTickets uint

    // ask user for their name, last name, email and number of tickets fmt.Println("Enter your first name: ")
    fmt.Println("Enter your first name: ")
    fmt.Scan(&firstName)

    fmt.Println("Enter your last name: ")
    fmt.Scan(&lastName)

    fmt.Println("Enter your email: ")
    fmt.Scan(&email)

    fmt.Println("Enter number of tickets: ")
    fmt.Scan(&userTickets)

    return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
    remainingTickets = remainingTickets - userTickets

    // create a map for usebor
    var userData = UserData {
        firstName: firstName,
        lastName: lastName,
        email: email,
        numberOfTickets: userTickets,
    }

    bookings = append(bookings, userData)

    fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
    fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
