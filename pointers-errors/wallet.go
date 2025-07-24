package pointerserrors

import "fmt"

type Bitcoin int 

type Wallet struct {
	balance Bitcoin
}

/* 
	Using pointers: 
	w *Wallet now becomes pointer to struct Wallet
*/ 


func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance =+ amount
}

func (w *Wallet) Balance() Bitcoin {
	fmt.Printf("address of balance in test is %p \n", &w.balance)
	return w.balance
}

