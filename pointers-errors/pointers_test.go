package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		//fmt.Printf("address of balance in test is %p \n", &wallet.balance)
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}
	assertError := func(t testing.TB, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("wanted an error but didn't get one")
		}
		if got.Error() != want.Error() {
			t.Errorf("got %q, want %q", got, want)
		}
	}
	assertNoError := func(t testing.TB, err error) {
		t.Helper()
		if err != nil {
			t.Fatal("got an error but didn't want one")
		}
	}
	t.Run("Deposit: ", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw: ", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(5))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(5))
	})
	t.Run("Withdraw insufficient funds: ", func(t *testing.T) {
		startingBalance := Bitcoin(30)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)

	})

}
