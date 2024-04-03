package main

import "fmt"

/*
Profit Calulator

Take the following as user input

revenue, expenses and tax rate

Calculate the earings before tax (EBT)

Calculate the earings after tax (EBT)

output EBT, profit and the ratio
*/

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Println("Please Enter the following information")
	fmt.Print("Total Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("\nTotal Expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("\nYour Current Tax Rate: ")
	fmt.Scan(&taxRate)

	earingsBeforeTax := revenue - expenses
	profit := earingsBeforeTax * (1 - taxRate/100)
	ratio := earingsBeforeTax / profit
	// use a printf to print them on the same line works the same as in c
	fmt.Printf("EBT: %f Profit: %f Ratio: %f \n", earingsBeforeTax, profit, ratio)

}
