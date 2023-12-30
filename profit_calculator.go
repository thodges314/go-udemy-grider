package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	tax := ebt * taxRate / 100
	profits := ebt - tax
	ratio := ebt / profits

	fmt.Printf("EBT: %.2f\n", ebt)
	fmt.Printf("Profits: %.2f\n", profits)
	fmt.Printf("Ratio: %.2f\n", ratio)

}
