package wallet

import (
	"errors"
	"fmt"
)

type Bitcoint int

type Stringer interface {
	String() string
}

func (b Bitcoint) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoint
}

func (w *Wallet) Withdraw(amount Bitcoint) error {
	if w.balance-amount < 0 {
		return errors.New("cannot withdraw, insufficient funds")
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Deposit(amount Bitcoint) {
	w.balance += amount
}

func (w Wallet) Balance() Bitcoint {
	return w.balance
}
