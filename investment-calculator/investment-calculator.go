// main entrypoint
package main

// imports
import (
	"fmt"
	"math"
)

// move const to be global
const inflationRate = 7.5 // you would prob just pass this in but for example of global scope constants where constants are fine to be global scope

// main function
func main() {

	// var investmentAmount float64 = 1000
	// var expectedReturnRate = 5.5
	// var years float64 = 10

	// short hand for declaring vars that infers the type and multi vars assigments
	//investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5

	// investmentAmount := 0.0
	// years := 0.0
	// expectedReturnRate := 0.0
	// can be empty as initlised to 0.0 by default for float
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	// take user input
	fmt.Print("Please input your Investment Amount: ")
	fmt.Scan(&investmentAmount) // sacn will overwrite the value initilised
	fmt.Print("\nPlease add your Investment Term in Years: ")
	fmt.Scan(&years)
	fmt.Print("\nPlease add your expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	// you can cast your vars using the following syntax
	// var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	// better way to write this would be to use the correct types at the start
	//futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	//futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// same as above using a functions example
	futureValue, futureRealValue := calculateFutureValues(investmentAmount, years, expectedReturnRate)
	// create a formatted string using sprintf
	formattedFv := fmt.Sprintf("Future Value: %.2f \n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (Adjusted for inflation) %.2f \n", futureRealValue)

	// output formatted
	//fmt.Printf("Future Value: %.2f \nFuture Value (Adjusted for inflation) %.2f \n", futureValue, futureRealValue)
	// output line
	//fmt.Println("Future Value:", futureValue)
	//fmt.Println(futureRealValue)
	fmt.Print(formattedFv, formattedRFV)
	// output using multiline printf
	/*
		fmt.Printf(`Future Value: %.2f
		Future Value (Adjusted for inflation) %.2f`, futureValue, futureRealValue)
	*/
}

// get future value function
func calculateFutureValues(investmentAmount float64, years float64, expectedReturnRate float64) (fv float64, rfv float64) {
	// keep in mind when returning values single returns only need type without () float64 multi returns uses (float64, int, string)
	// if you declare your return types you can include the var names and they dont need to be initilise in the function
	// if you do not set the var name in the return value you should initilise them then return the and only use the type turn in the function
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)

	return fv, rfv
}
