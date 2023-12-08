package main

import (
	"banksy/fileops"
	"fmt"
	"time"
)

const defaultBalance float64 = 1000

type account struct {
	balance float64
	io      *fileops.FileIO
}

func main() {
	resetScreen()
	fmt.Println("Welcome to Banksy!")
	stream := fileops.CreateFileStream("balance.txt")
	a := account{
		balance: 0,
		io:      stream,
	}
	accountBalance, err := a.getBalance(defaultBalance)
	if err != nil {
		fmt.Println(err)
	}
	a.balance = accountBalance
	time.Sleep(time.Second * 3)
	resetScreen()
	a.presentOptions()
}

func (a *account) saveBalance() {
	a.io.WriteFloatToFile(a.balance)
}

func (a *account) getBalance(init float64) (float64, error) {
	return a.io.GetFloatFromFile(init)
}

func (a *account) checkBalance() {
	resetScreen()
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
	resetScreen()
	fmt.Printf("Deposited %.2f into your account\n", amount)
	fmt.Printf("Your new balance: %v\n", balance)
	a.saveBalance()
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
	resetScreen()
	if a.balance-amount < 0 {
		fmt.Println("Sorry! I can't do that, you don't have enough funds.")
	} else {
		a.balance -= amount
		balance := fmt.Sprintf("%.2f", a.balance)
		fmt.Printf("Withdrew %.2f from your account\n", amount)
		fmt.Printf("Your new balance: %v\n", balance)
		a.saveBalance()
	}
}
