package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()

		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s, wanted %s", got, want)
		}
	}

	assertError := func(t testing.TB, got error, want string) {
		t.Helper()
		if got == nil {
			//ends without running any following assertions
			t.Fatal("dind't get an error but wanted one.")
		}

		if got.Error() != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}

	assertNotError := func (t testing.TB, err error)  {
		t.Helper()

		if err != nil {
			t.Fatal("got an error but didn't expected one")
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)

		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(100)}

		err := wallet.Withdraw(Bitcoin(10))

		want := Bitcoin(90)
		
		assertNotError(t, err)

		assertBalance(t, wallet, want)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(10)

		wallet := Wallet{balance: startingBalance}

		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)

		assertError(t, err, ErrInsuficientFunds.Error())
	})

}
