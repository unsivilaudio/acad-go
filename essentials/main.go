package main

import (
	"fmt"
	PC "profit_calculator"
)

func main() {
	fmt.Printf("Started main program!\n")
	// IC.RunInvestment()
	calc := PC.GenerateCalculator()
	calc.CalculateProfit()
}
