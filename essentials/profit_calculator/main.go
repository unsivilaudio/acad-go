package profit_calculator

import (
	"errors"
	"fmt"
	"log"
)

type calculator struct {
	revenue  float64
	expenses float64
	taxRate  float64
}

func GenerateCalculator() calculator {
	var revenue float64
	var expenses float64
	var taxRate float64

	err := getUserPrompt(&revenue, "total revenue")
	if err != nil {
		log.Fatalf("%v", err)
	}
	err = getUserPrompt(&expenses, "total expenses")
	if err != nil {
		log.Fatalf("%v", err)
	}
	err = getUserPrompt(&taxRate, "tax rate")
	if err != nil {
		log.Fatalf("%v", err)
	}
	calc := calculator{
		revenue:  revenue,
		expenses: expenses,
		taxRate:  taxRate,
	}

	return calc
}

func getUserPrompt(field *float64, label string) error {
	fmt.Printf("Enter your %s: ", label)
	fmt.Scan(field)
	if field == nil {
		return errors.New("You must provide a valid value for revenue.")
	}

	return nil
}

func (c *calculator) CalculateProfit() {
	intermediateProfit := c.revenue - c.expenses
	collectedTax := c.revenue * c.taxRate / 100
	totalProfit := intermediateProfit - collectedTax

	fmt.Printf("\nIntermediate profit: %.2f", intermediateProfit)
	fmt.Printf("\nTax collected: %.2f", collectedTax)
	fmt.Printf("\nTotal profit: %.2f", totalProfit)
}
