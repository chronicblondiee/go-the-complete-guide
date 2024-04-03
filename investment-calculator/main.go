// main entrypoint
package main

// imports
import (
	"fmt"
	"math"
)

// main function
func main() {

	var investmentAmount float64 = 1000
	var expectedReturnRate = 5.5
	var years float64 = 10

	// you can cast your vars using the following syntax
	// var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	// better way to write this would be to use the correct types at the start
	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	//fmt.Printf("%.2f \n", futureValue)
	fmt.Println(futureValue)

}
