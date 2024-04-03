// main entrypoint
package main

// imports
import (
	"fmt"
	"math"
)

// main function
func main() {

	// var investmentAmount float64 = 1000
	// var expectedReturnRate = 5.5
	// var years float64 = 10

	// short hand for declaring vars that infers the type and multi vars assigments
	//investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5
	const inflationRate = 7.5
	investmentAmount := 0.0
	years := 10.0
	expectedReturnRate := 5.5

	// take user input
	fmt.Print("Please input your Investment Amount: ")
	fmt.Scan(&investmentAmount) // sacn will overwrite the value initilised
	fmt.Print("\nPlease add your Investment Term: ")
	fmt.Scan(&years)
	fmt.Print("\nPlease add your expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	// you can cast your vars using the following syntax
	// var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	// better way to write this would be to use the correct types at the start
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	//fmt.Printf("%.2f \n", futureValue)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

}
