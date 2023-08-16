package main

import (
	"fmt"
)

//go mod init goApp to generate the mod file

func main() {
	//fmt.Print("hello world")
	var totalTickets = 50
	var remainingTickets uint = 50
	var firstname string
	var lastname string
	var noOfTickets uint
	// booking1 := []string{}
	var bookings []string

	for {
		if remainingTickets <= 0 {
			fmt.Print("bookings are closed")
			break
		}
		fmt.Print("Enter you first name:")
		fmt.Scan(&firstname)

		fmt.Print("Enter you last name:")
		fmt.Scan(&lastname)

		fmt.Print("Enter No of tickets:")
		fmt.Scan(&noOfTickets)

		if noOfTickets > remainingTickets {
			fmt.Printf("we have only %v tickets remaining,so cant book %v tickets\n", remainingTickets, noOfTickets)
			break
		}

		bookings = append(bookings, firstname+" "+lastname)
		fmt.Printf("These all are our bookings%v", bookings)
		fmt.Println()

		remainingTickets = remainingTickets - noOfTickets
		fmt.Printf("Total tickets are %v \nThe remaining tickets are %v\n", totalTickets, remainingTickets)
		lastnames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			var lastname1 = names[1]
			lastnames = append(lastnames, lastname1)
			fmt.Printf("names of bookings are%v\n", lastnames)
		}
	}
	var city string
	for i := 0; i < 5; i++ {
		fmt.Print("Enter you city:")
		fmt.Scan(&city)
		switch city {
		case "vzm":
			fmt.Printf("you city is %v\n", city)
		case "vsp":
			fmt.Printf("you city is %v\n", city)
		case "rbp":
			fmt.Printf("you city is %v\n", city)
		default:
			fmt.Print("your city is vzmm\n")
		}

	}

}
