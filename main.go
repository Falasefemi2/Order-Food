package main

import (
	"fmt"
	"io"
	"os"
)

// make subTotalBill as global variable to make it easily accessible in case customer modifies her order.
// subTotalBill means total bill but excluding taxes.
var subTotalBill float64

// make a map of customerOrder in which we will store the items ordered as "key" and no. of plates as "value".
var customerOrder = make(map[string]uint, 0)

func main() {
	var customerName string
	fmt.Println("What is your first name?")
	fmt.Scan(&customerName)

	greet(customerName)
	orderItems()
	displayGeneratingBill() //just displays that "generating bill" in a fancy manner.
	generateBill()
	modifyOrder()
	// Create a temporary file
	tempFile, err := os.CreateTemp("", "bill*.txt")
	if err != nil {
		fmt.Println("Error creating temp file:", err)
		return
	}
	defer os.Remove(tempFile.Name()) // Clean up the temp file

	// Redirect standard output to the temp file
	old := os.Stdout
	defer func() {
		os.Stdout = old // Restore the original stdout
	}()
	os.Stdout = tempFile

	printFinalBill()
	// Read the contents of the temp file
	tempFile.Seek(0, 0) // Reset the file pointer
	billBytes, err := io.ReadAll(tempFile)
	if err != nil {
		fmt.Println("Error reading temp file:", err)
		return
	}
	billString := string(billBytes)

	// Save the bill to the final file
	saveBillToFile(billString, customerName)

	sayTata(customerName)
}

// Function to save the bill to a file
// func saveBillToFile(bill string) {
// 	fileName := "final_bill.txt"
// 	file, err := os.Create(fileName)
// 	if err != nil {
// 		fmt.Println("Error creating file:", err)
// 		return
// 	}
// 	defer file.Close()

// 	_, err = file.WriteString(bill)
// 	if err != nil {
// 		fmt.Println("Error writing to file:", err)
// 		return
// 	}

// 	fmt.Println("Bill saved to", fileName)
// }

// Function to save the bill to a file
func saveBillToFile(bill string, customerName string) {
	fileName := customerName + "_bill.txt"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(bill)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Printf("Bill saved to %s\n", fileName)
}
