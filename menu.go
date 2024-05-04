package main

import (
	"fmt"
	"strings"
)

type Menu struct {
	itemNo    uint
	itemName  string
	itemPrice float64
}

var menu = []Menu{
	// itemNo itemName itemPrice
	{1, "Jollof Rice", 3000},
	{2, "Coffee", 1500},
	{3, "Amala", 2500},
	{4, "Eba", 700},
	{5, "Egusi Soup", 300},
	{6, "Ice Cream", 600},
	{7, "Egg Roll", 300},
	{9, "Foreign Dishes", 15000},
	{10, "Pounded Yam", 300},
	{11, "Yam and Egg", 5500},
	{12, "Salad", 150},
	{13, "Semo", 500},
}

func printMenu() {
	fmt.Printf("%15s\n", "Menu")
	fmt.Printf("%s\n", strings.Repeat("-", 35))
	fmt.Printf("%-7s %6s  %12s\n", "S.No", "Price", "Item Name")
	fmt.Printf("%s\n", strings.Repeat("-", 35))

	for _, element := range menu {
		fmt.Printf(" %-7d %.2f   %-4s\n", element.itemNo, element.itemPrice, element.itemName)
	}
	fmt.Printf("%s", strings.Repeat("-", 35))
	fmt.Println()
}
