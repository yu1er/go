package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, 10)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, 10)
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		initBalance := Bitcoin(20)
		wallet := Wallet{balance: initBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, initBalance)
		assertError(t, err, ErrorInsufficientFunds)
	})
}

func assertNoError(t *testing.T, got error) {
	if got != nil {
		t.Fatal("got an error but didnt want one")
	}

}

func assertError(t *testing.T, got error, want error) {
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

func assertBalance(t *testing.T, wallet Wallet, expected Bitcoin) {
	got := wallet.Balance()

	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}
}
