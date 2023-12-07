package investment_calculator

import (
	"errors"
	"fmt"
	"log"
	"math"
)

type calculator struct {
	investmentAmount int
	expectedReturn   float64
	years            int
}

func GenerateCalculator(invested int, interestRate float64, duration int) (error, calculator) {
	calc := calculator{}
	if invested <= 0 {
		err := "You must provide a valid investment value"
		return errors.New(err), calc
	}
	if interestRate <= 0 {
		err := "You must provide a valid interest rate value"
		return errors.New(err), calc
	}
	if duration <= 0 {
		err := "You must provide a valid duration (years) value"
		return errors.New(err), calc
	}

	calc.investmentAmount = invested
	calc.expectedReturn = interestRate
	calc.years = duration

	return nil, calc
}

func (c *calculator) CalculateAndPrint() {
	inflationRate := 2.5
	futureReturn := float64(c.investmentAmount) * math.Pow(1+c.expectedReturn/100, float64(c.years))
	futureRealReturn := float64(c.investmentAmount) * math.Pow(1+inflationRate/100, float64(c.years))

	fmt.Printf("\nInitial investment amount: %v", c.investmentAmount)
	fmt.Printf("\nInterest Rate: %v", c.expectedReturn)
	fmt.Printf("\nDuration (years): %v", c.years)
	fmt.Println("\n================")
	fmt.Printf("Accrued Value: +%v", math.Round(futureReturn-float64(c.investmentAmount)))
	fmt.Printf("\nTotal Value: %v", math.Round(futureReturn))
	fmt.Printf("\nInflation Adjusted Value: %v", math.Round(futureRealReturn))
}

func RunInvestment() {
	var investAmount int
	var investRate float64
	var investDuration int

	fmt.Print("Enter amount to be invested: ")
	fmt.Scan(&investAmount)
	fmt.Print("Enter target interest rate: ")
	fmt.Scan(&investRate)
	fmt.Print("Enter investment duration (years): ")
	fmt.Scan(&investDuration)

	err, calc := GenerateCalculator(investAmount, investRate, investDuration)
	if err != nil {
		log.Fatalf("%v", err)
	}

	calc.CalculateAndPrint()
}
