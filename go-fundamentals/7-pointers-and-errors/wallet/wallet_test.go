package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("desosit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoint(10))

		got := wallet.Balance()
		want := Bitcoint(10)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("desosit", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoint(20)}
		wallet.Withdraw(Bitcoint(10))

		got := wallet.Balance()
		want := Bitcoint(10)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
