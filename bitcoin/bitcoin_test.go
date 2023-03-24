package bitcoin

import "testing"

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		showBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		showBalance(t, wallet, Bitcoin(10))
		notFindError(t, err, nil)
	})

	t.Run("Insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))
		showBalance(t, wallet, Bitcoin(20))
		findError(t, err, WithoutBalance)
	})
}

/* Auxiliary functions */
func showBalance(t *testing.T, w Wallet, expected Bitcoin) {
	t.Helper()
	result := w.Balance()
	if result != expected {
		t.Errorf("result %s, expected %s", result, expected)
	} 
}

func findError(t *testing.T, err error, expected error) {
	t.Helper()
	if err == nil {
		t.Fatal("expected an error")
	}
	if err != expected {
		t.Errorf("result %s, expected %s", err, expected)
	}
}

func notFindError(t *testing.T, err error, expected error) {
	t.Helper()
	if err != nil {
		t.Fatal("expected an error")
	}
}