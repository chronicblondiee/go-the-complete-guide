package main

import "fmt"

/*
Profit Calulator

Take the following as user input

revenue, expenses and tax rate

Calculate the earings before tax (EBT)

Calculate the earings after tax (Profit)

output EBT, profit and the ratio of EBT and Profit
*/

func main() {

	revenue, expenses, taxRate := takeInput()

	earingsBeforeTax, profit, ratio := doTaxCalculations(revenue, expenses, taxRate)
	// use a printf to print them on the same line works the same as in c
	fmt.Printf("EBT: %.2f Profit: %.2f Ratio: %.2f \n", earingsBeforeTax, profit, ratio)

}

func takeInput() (revenue float64, expenses float64, taxRate float64) {
	// take user input
	fmt.Println("Please Enter the following information")
	fmt.Print("Total Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("\nTotal Expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("\nYour Current Tax Rate: ")
	fmt.Scan(&taxRate)

	return revenue, expenses, taxRate
}

func doTaxCalculations(revenue float64, expenses float64, taxRate float64) (earingsBeforeTax float64, profit float64, ratio float64) {

	earingsBeforeTax = revenue - expenses
	profit = earingsBeforeTax * (1 - taxRate/100)
	ratio = earingsBeforeTax / profit

	return earingsBeforeTax, profit, ratio
}
