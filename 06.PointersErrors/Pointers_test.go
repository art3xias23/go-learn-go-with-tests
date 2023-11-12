package Pointers

import "testing"

func Test_Wallet(t *testing.T) {
	t.Run("Test deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10.0)
		want := Bitcoin(10.0)
		showError(t, wallet, want)
	})

	assertError:=func(t *testing.T, err error){
		if err == nil{
			t.Error("wanted an errlr but didn't get one")
		}
	}

	t.Run("Test withdrawal", func(t *testing.T) {
		w := Wallet{}
		w.Deposit(Bitcoin(20.0))
		w.Withdrawal(Bitcoin(5.0))
		want := Bitcoin(15.0)

		showError(t, w, want)
	})

	t.Run("withdrawal insufficient funds", func(t *testing.T){
		startingBalance := Bitcoin(20)
		wallet:=Wallet{startingBalance}
		err := wallet.Withdrawal(100)
		showError(t, wallet, startingBalance)
		assertError(t, err)
	})
}

func showError(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()

	got:=wallet.Balance()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}
