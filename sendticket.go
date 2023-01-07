package main

import (
	"fmt"
	"time"
)

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("################")
	fmt.Printf("Sending ticket: \n %v \n to email address %v\n", ticket, email)
	fmt.Println("#################")

	wg.Done() // the thread is done
}
