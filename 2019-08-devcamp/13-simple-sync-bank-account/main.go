package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	balance float64
	mu      sync.Mutex
}

func (b *BankAccount) Deposit(amount float64) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.balance += amount
}

func (b *BankAccount) Withdraw(amount float64) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.balance -= amount
}

func (b *BankAccount) Balance() float64 {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.balance
}

func main() {

	var wg sync.WaitGroup

	acc := BankAccount{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(account *BankAccount) {
			account.Deposit(10)
			account.Withdraw(account.Balance() / 2)
			wg.Done()
		}(&acc)
	}
	wg.Wait()

	fmt.Println(acc.Balance())

}
