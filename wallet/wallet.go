package main

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	//any variable or function name start with a small letter is private opposite is public
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}