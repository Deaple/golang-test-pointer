package main

import (
	"testing"
)


func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()

	// var want Bitcoin
	// want = 10
	want := Bitcoin(20)

	if got != want {
		t.Errorf("got %s, wanted %s",got,want)
	}


}