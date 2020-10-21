package main

type Customer struct {
	creditCard CreditCard
}

type CreditCard struct {
	number string
}

func (c *Customer) PayWithCreditCard() {
	//... first attempt
	logPaymentTransaction()
	fallbackToSecondAcquirer()
}

func (c *Customer) fallbackToSecondAcquirer() {
	//... logic here
	logPaymentTransaction()
}

func (c *Customer) logPaymentTransaction() {
	//... logic here
}

ty