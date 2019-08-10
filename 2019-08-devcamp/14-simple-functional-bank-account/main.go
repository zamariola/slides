package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	balance    float64
	operations chan func(float64) float64
}

func (b *BankAccount) Deposit(amount float64) {
	b.operations <- func(oldBalance float64) float64 {
		return oldBalance + amount
	}
}

func (b *BankAccount) Withdraw(amount float64) {
	b.operations <- func(oldBalance float64) float64 {
		return oldBalance - amount
	}
}

func (b *BankAccount) Balance() float64 {
	return b.balance
}

func (b *BankAccount) Loop() {
	for op := range b.operations {
		b.balance = op(b.balance)
	}
}

func NewBankAccount() *BankAccount {
	return &BankAccount{
		balance:    0,
		operations: make(chan func(float64) float64),
	}
}

func main() {

	var wg sync.WaitGroup

	acc := NewBankAccount()

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(account *BankAccount) {
			account.Deposit(10)
			account.Withdraw(account.Balance() / 2)
			wg.Done()
		}(acc)
	}

	go func(account *BankAccount) {
		acc.Loop()
	}(acc)
	wg.Wait()

	fmt.Println(acc.Balance())

}
