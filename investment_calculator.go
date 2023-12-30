package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount, years, expectedReturnRate float64 = 1000, 10, 5.5

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
