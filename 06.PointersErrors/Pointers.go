package Pointers

import "fmt"

type (
	Bitcoin int
	Wallet  struct {
		balance Bitcoin
	}
)

func (w *Wallet) Deposit(deposit Bitcoin) {
	w.balance += deposit 
}

func (w Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d Bitcoin", b)
}
