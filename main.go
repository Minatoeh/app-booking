package main

import (
	"app-booking/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	firstName, lastName, userTickets, userMail := getUserInput()

	// Pass remainingTickets to ValidateUserInput
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(userTickets, firstName, lastName, userMail, remainingTickets)
	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(userTickets, firstName, lastName, userMail)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, userMail)
		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			// end program
			fmt.Println("All tickets are sold out. We're Sorry!")
			//break
		}
	} else {
		if !isValidName {
			fmt.Println("Your name is invalid. Please provide valid information.")
		}
		if !isValidEmail {
			fmt.Println("Your email information is wrong.")
		}
		if !isValidTicketNumber {
			fmt.Println("It's not possible to buy this amount of tickets.")
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to our %v Conference\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("It's time to check your tickets.")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		//var names = strings.Fields(booking)
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, uint, string) {
	var firstName string
	var lastName string
	var userTickets uint
	var userMail string

	fmt.Println("Tell us your first name please:")
	fmt.Scan(&firstName)

	fmt.Println("What's your surname:")
	fmt.Scan(&lastName)

	fmt.Println("How many tickets do you need:")
	fmt.Scan(&userTickets)

	fmt.Println("Tell us your email:")
	fmt.Scan(&userMail)

	return firstName, lastName, userTickets, userMail
}

func bookTicket(userTickets uint, firstName string, lastName string, userMail string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           userMail,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("User %v %v booked %v tickets. User e-mail is: %v.\n", firstName, lastName, userTickets, userMail)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, userMail string) {
	//fmt.Println("Wait for a moment please. We're prepared your invitation...")
	time.Sleep(25 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#######")
	fmt.Printf("Sending ticket via\n %v to email adress %v\n", ticket, userMail)
	fmt.Println("#######")
	wg.Done()
}
