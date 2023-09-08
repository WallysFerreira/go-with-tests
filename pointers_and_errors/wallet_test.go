package main

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()
	expected := Bitcoin(10)

	if got != expected {
		t.Errorf("%#v got %d expected %d", wallet, got, expected)
	}
}
