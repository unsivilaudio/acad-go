package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/inancgumus/screen"
)

const defaultBalance = 1000.00

type account struct {
	balance float64
}

func main() {
	fmt.Println("Welcome to Banksy!")
	a := account{
		balance: defaultBalance,
	}
	accountBalance, err := a.getBalanceFromFile()
	if err != nil {
		a.balance = accountBalance
		fmt.Println(err)
	}

	time.Sleep(time.Second * 3)
	a.getUserPrompt()
}

func (a *account) writeBalanceToFile() {
	balanceText := fmt.Sprintf("%.2f", a.balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func (a *account) getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile("balance.txt")
	if err != nil {
		return defaultBalance, errors.New("Failed to find balance file.")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return defaultBalance, errors.New("Failed to parse stored balance value.")
	}
	return balance, nil
}

func (a *account) getUserPrompt() {
	var choice int
	screen.Clear()
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	fmt.Scan(&choice)

	switch os := choice; os {
	case 1:
		a.checkBalance()
		time.Sleep(time.Second * 3)
		a.getUserPrompt()
		break
	case 2:
		a.makeDeposit()
		time.Sleep(time.Second * 3)
		a.getUserPrompt()
		break
	case 3:
		a.makeWithdrawal()
		time.Sleep(time.Second * 3)
		a.getUserPrompt()
		break
	case 4:
		fmt.Println("Thanks for banking with us today. Bye!")
		break
	default:
		a.getUserPrompt()
	}
}

func (a *account) checkBalance() {
	balance := fmt.Sprintf("%.2f", a.balance)
	fmt.Printf("\nYour balance: %v", balance)
}

func (a *account) makeDeposit() {
	var amount float64
	fmt.Printf("Please enter an amount: ")
	fmt.Scan(&amount)
	if amount <= 0 {
		fmt.Println("Please enter a valid amount.")
		a.makeDeposit()
		return
	}
	a.balance += amount
	balance := fmt.Sprintf("%.2f", a.balance)
	fmt.Printf("Deposited %.2f into your account\n", amount)
	fmt.Printf("Your new balance: %v\n", balance)
	a.writeBalanceToFile()
}

func (a *account) makeWithdrawal() {
	var amount float64
	fmt.Printf("Please enter an amount: ")
	fmt.Scan(&amount)
	if amount <= 0 {
		fmt.Println("Please enter a valid amount.")
		a.makeWithdrawal()
		return
	}
	if a.balance-amount < 0 {
		fmt.Println("Sorry! I can't do that, you don't have enough funds.")
	} else {
		a.balance -= amount
		balance := fmt.Sprintf("%.2f", a.balance)
		fmt.Printf("Withdrew %.2f from your account\n", amount)
		fmt.Printf("Your new balance: %v\n", balance)
		a.writeBalanceToFile()
	}
}
