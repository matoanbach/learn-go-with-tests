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

	assertError := func(t testing.TB, err error) {
		t.Helper()
		if err == nil {
			t.Errorf("wanted an error but didn't get one")
		}
	}

	t.Run("desosit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoint(10))

		want := Bitcoint(10)
		assertBalance(t, wallet, want)
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoint(20)}
		wallet.Withdraw(Bitcoint(10))

		want := Bitcoint(10)
		assertBalance(t, wallet, want)
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoint(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoint(100))

		assertError(t, err)
		assertBalance(t, wallet, startingBalance)
	})
}
