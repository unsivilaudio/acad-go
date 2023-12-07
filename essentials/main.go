package main

import (
	"fmt"
	IC "investment_calculator"
	"log"
)

func main() {
	fmt.Printf("This is the main function!")
	err, calc := IC.GenerateCalculator(1500, 7.5, 7)
	if err != nil {
		log.Fatalf("%v", err)
	}

	calc.CalculateAndPrint()
}
