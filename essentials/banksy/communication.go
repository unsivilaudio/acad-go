package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func (a *account) presentOptions() {
	var choice int
	resetScreen()
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	fmt.Scan(&choice)

	switch os := choice; os {
	case 1:
		a.checkBalance()
	case 2:
		a.makeDeposit()
	case 3:
		a.makeWithdrawal()
	case 4:
		fmt.Println("Thanks for banking with us today. Bye!")
		return
	default:
		a.presentOptions()
	}
	time.Sleep(time.Second * 3)
	a.presentOptions()
}

func resetScreen() {
	screen.Clear()
	screen.MoveTopLeft()
}
