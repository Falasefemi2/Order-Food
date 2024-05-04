package main

import "fmt"

func greet(customerName string) {
	fmt.Printf("%52s %s\n", "Greetings ", customerName)
	fmt.Printf("%72s\n", "_/\\_ Welcome to Falase Femi! _/\\_ ")
	fmt.Println()
}

func sayTata(customerName string) {
	fmt.Println()
	fmt.Printf("%17s", " ")
	fmt.Printf("_/\\_ Thank you %v for visiting Falase Femi! _/\\_\n", customerName)
	fmt.Printf("%55s\n", "We hope to see you again!")
	fmt.Printf("%51s\n", "Havae a nice day! ")
}
