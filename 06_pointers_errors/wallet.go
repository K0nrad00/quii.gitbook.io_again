package main

import (
	"errors"
	"fmt"
)

type Wallet struct {
	balance Bitcoin
}

//Remember we can access the internal/private 'balance' field in the struct using the "receiver" variable.
func (w *Wallet) Deposit(amount Bitcoin) {
	/*
		You can see that the addresses of the two balances are different. So when we change the value of the balance inside the code, we are working on a copy of what came from the test. Therefore the balance in the test is unchanged.

		I THINK THIS IS HAPPENING AS 'balance' IS PRIVATE, IN PREVIOUS EXAMPLES WE HAD NO PRIVATE VARIABLES - THEY WERE ACCESSIBLE BY TEST AT THE RUNTIME..
		DOES THIS MEAN THAT ANY PRIVATE VARIABLE WOULD NEED TO BE ACCESSED BY POINTER OUTSIDE THAT FILE ?
	*/
	// fmt.Printf("NOT_TEST: address of balance: %v, value: %v\n", &w.balance, w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
