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

	assertError := func(t testing.TB, got error, want string) {
		t.Helper()
		if got == nil {
			t.Errorf("wanted an error but didn't get one")
		}

		if got.Error() != want {
			t.Errorf("got %q, want %q", got, want)
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
		want := "cannot withdraw, insufficient funds"
		err := wallet.Withdraw(Bitcoint(100))

		assertError(t, err, want)
		assertBalance(t, wallet, startingBalance)
	})
}
