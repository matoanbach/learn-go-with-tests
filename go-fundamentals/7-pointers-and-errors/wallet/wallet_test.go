package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoint) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}
	t.Run("desosit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoint(10))

		want := Bitcoint(10)
		assertBalance(t, wallet, want)
	})
	t.Run("desosit", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoint(20)}
		wallet.Withdraw(Bitcoint(10))

		want := Bitcoint(10)
		assertBalance(t, wallet, want)
	})
}
