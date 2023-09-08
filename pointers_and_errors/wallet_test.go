package main

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()
	expected := 10

	if got != expected {
		t.Errorf("%#v got %d expected %d", wallet, got, expected)
	}
}
