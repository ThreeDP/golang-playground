package wallet

type Wallet struct {
	balance 	int
}
/*
Quando func (c Carteira) Depositar(quantidade int) é chamado,
o c é uma cópia do valor de qualquer lugar que o método
tenha sido chamado. */
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
}