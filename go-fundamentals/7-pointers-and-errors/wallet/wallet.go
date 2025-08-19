package wallet

import "fmt"

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

func (w *Wallet) Withdraw(bitcoint Bitcoint) {
	w.balance -= bitcoint
}

func (w *Wallet) Deposit(amount Bitcoint) {
	w.balance += amount
}

func (w Wallet) Balance() Bitcoint {
	return w.balance
}
