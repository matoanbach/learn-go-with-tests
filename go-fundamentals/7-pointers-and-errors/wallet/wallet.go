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

func (w *Wallet) Deposit(amount Bitcoint) {
	fmt.Printf("address of my wallet in wallet.go is %p \n", w)
	w.balance += amount
}

func (w Wallet) Balance() Bitcoint {
	return w.balance
}
