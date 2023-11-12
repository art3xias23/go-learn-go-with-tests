package Pointers

import "testing"

func Test_Wallet(t *testing.T) {
	t.Run("Test deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10.0)
		got := wallet.Balance()
		want := Bitcoin(10.0)
		showError(t, want, got)
	})

	t.Run("Test withdrawal", func(t *testing.T) {
		w := Wallet{}
		w.Deposit(Bitcoin(20.0))
		w.Withdrawal(Bitcoin(5.0))
		got := w.Balance()
		want := Bitcoin(15.0)

		showError(t, want, got)
	})
}

func showError(t *testing.T, want Bitcoin, got Bitcoin) {
	t.Helper()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}
