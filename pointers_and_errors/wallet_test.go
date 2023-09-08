package main

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		expected := Bitcoin(10)

		if got != expected {
			t.Errorf("%#v got %s expected %s", wallet, got, expected)
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()
		expected := Bitcoin(10)

		if got != expected {
			t.Errorf("%#v got %s expected %s", wallet, got, expected)
		}
	})

}
