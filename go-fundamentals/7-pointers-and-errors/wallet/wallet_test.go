package wallet

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)

	got := wallet.Balance()
	want := Bitcoint(10)
	fmt.Printf("address of my wallet in test is %p \n", &wallet)
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
