package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func displayGeneratingBill() {
	fmt.Println()
	billDisplayTest := "************************************* Generating Bill *************************************"
	for _, element := range billDisplayTest {
		fmt.Printf("%c", element)
		time.Sleep(time.Millisecond * 15)
	}
}

// prints item name, price, quantity and total price and sub total amount.
func generateBill() {
	fmt.Println()
	fmt.Printf("+%s=\n", strings.Repeat("-", 90))
	fmt.Printf(" %-30s %-20s %-20s\n", "Item Name", "Price($)", "Total Price($)")
	fmt.Printf("+%s+\n", strings.Repeat("-", 90))

	// prints the details of the food item that you'ce ordered.
	printOrderData()

	fmt.Printf("+%s+\n", strings.Repeat("-", 90))

	// print sub total amount in a cool way!
	fmt.Printf("%47s", " ")
	for _, element := range "Sub Total(excluding GST): $" {
		fmt.Printf("%c", element)
		time.Sleep(time.Millisecond * 50)
	}
	fmt.Printf("%.2f\n", subTotalBill)
}

// prints the data of the items that you  ordered
func printOrderData() {
	for key := range customerOrder {
		//key -> it is the key values
		for _, element := range menu {
			if key == element.itemName {
				//customerOrder[key] -> will provide the no. of plates of that item
				//float64(customerOrder[key])*element.itemPrice -> this will result in the cost of each item
				totalCostOfItem := float64(customerOrder[key]) * element.itemPrice
				fmt.Printf(" %-30s %-20.2f %-20v %-20.2f\n", key, element.itemPrice, customerOrder[key], totalCostOfItem)
			}
		}
	}
	fmt.Println()
}

func printFinalBill() {
	for _, element := range "Here is your fianl bill:-" {
		fmt.Printf("%c", element)
		time.Sleep(time.Millisecond * 50)
	}
	fmt.Println()

	fmt.Printf("\n%52s\n", "FALASE FEMI")
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%s\n", strings.Repeat("*", 91))
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%86s\n", "37 Oba Ni Jesu Road Ijoko, Ogun State")
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%50s\n", "TEL: +2347013329953")
	fmt.Printf("%60s\n\n", "Email: femifalase228@gmail.com")
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%s", strings.Repeat("-", 42))
	fmt.Printf("%s", "INVOICE")
	fmt.Printf("%s\n", strings.Repeat("-", 42))
	time.Sleep(time.Millisecond * 200)

	// rand.Intn(100) //necessary to produce random integers
	rand.NewSource(time.Now().UnixNano())
	fmt.Printf(" Ticket No: %d\n", rand.Intn(550)+1)

	fmt.Printf(" Date: %v\n", time.Now().Local().Format("02-Jan-06")) //display date
	fmt.Printf(" Time: %v", time.Now().Local().Format("3:4 PM"))      //display time
	time.Sleep(time.Millisecond * 200)

	generateBill() //prints details of the bill,like, item name, price, quantity and total price and sub total amount.

	tax := 18 * subTotalBill / (100)
	grandTotal := subTotalBill + tax

	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%71s: $%.2f\n", "GST", tax) //display tax amount
	fmt.Printf("+%s+\n", strings.Repeat("-", 90))
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%71s: $%.2f\n", "Grand Total", grandTotal) //display final bill
	fmt.Printf("+%s+\n", strings.Repeat("-", 90))
}
