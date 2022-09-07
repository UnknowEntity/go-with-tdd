package wallet

import (
	"errors"
	"fmt"
)

type Stringer interface {
	String() string
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	ballance Bitcoin
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Deposit(amount Bitcoin) {
	w.ballance += amount
}

func (w *Wallet) Ballance() Bitcoin {
	return w.ballance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.ballance {
		return ErrInsufficientFunds
	}

	w.ballance -= amount
	return nil
}
