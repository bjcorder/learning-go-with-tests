package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (w *Wallet) Deposit(amount Bitcoin) {
	//fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	//fmt.Printf("address of balance in Withdrawal is %p \n", &w.balance)
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	fmt.Printf("address of balance in Balance is %p \n", &w.balance)
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
