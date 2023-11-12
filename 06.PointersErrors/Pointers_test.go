package Pointers

import "testing"

func Test_Wallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10.0)
	got := wallet.Balance()
	want := Bitcoin(10.0)
	showError(t, want, got)
}

func showError(t *testing.T, want Bitcoin, got Bitcoin) {
	t.Helper()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
