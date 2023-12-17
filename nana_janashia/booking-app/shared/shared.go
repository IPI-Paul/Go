package shared

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint,
	remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func SendTicket(userTickets uint, firstname string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstname, lastName)
	fmt.Println("________________________")
	fmt.Printf("Sending ticket: \n%v \nto email address %v\n", ticket, email)
	fmt.Println("________________________")
}

func SendTicketWait(userTickets uint, firstname string, lastName string, email string, wg *sync.WaitGroup) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstname, lastName)
	fmt.Println("________________________")
	fmt.Printf("Sending ticket: \n%v \nto email address %v\n", ticket, email)
	fmt.Println("________________________")
	defer wg.Done()
}
