package Pointers

import "fmt"
import "errors"

type (
	Bitcoin int
	Wallet  struct {
		balance Bitcoin
	}
)

func (w *Wallet) Deposit(deposit Bitcoin) {
	w.balance += deposit
}

func (w *Wallet) Withdrawal(b Bitcoin) error {
	if b> w.Balance() {
		return errors.New("oh no, balance not enough")
	}
	w.balance -= b
	return nil
}

func (w Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d Bitcoin", b)
}
