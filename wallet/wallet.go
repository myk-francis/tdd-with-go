package main

type Wallet struct {
}

func (w Wallet) Deposit(amount int) {
	//any variable or function name start with a small letter is private opposite is public
	w.balance += amount
}

func (w Wallet) Balance() int {
	return w.balance
}