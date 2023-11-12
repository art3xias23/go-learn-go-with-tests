package Pointers

type Bitcoin int
type Wallet struct {
	balance Bitcoin 
}

func (w *Wallet) Deposit(deposit Bitcoin) {
	w.balance += deposit
}

func (w Wallet) Balance() Bitcoin {
	return w.balance
}
