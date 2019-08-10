package main

import "fmt"

type BankAccount struct {
	balance float64
}

func (b *BankAccount) Deposit(amount float64) { b.balance += amount }

func (b *BankAccount) Withdraw(amount float64) { b.balance -= amount }

func (b *BankAccount) Balance() float64 { return b.balance }

func main() {

	acc := BankAccount{}
	acc.Deposit(50)
	acc.Withdraw(40)
	fmt.Println(acc.Balance())
}
