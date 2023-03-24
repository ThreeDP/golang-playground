package wallet

import "testing"

func TestWallet(t *testing.T) {
	w := Wallet{}
	w.Deposit(10)
	result := w.Balance()
	expected := 10
	if result != expected {
		t.Errorf("result %d, expected %d", result, expected)
	}
}