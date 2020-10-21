package main

type Customer struct {
	creditCard     CreditCard
	additionalcard CreditCard
	billingAddress Address
	preferences    string
}

type BusinessCustomer struct {
	Customer
}

type CreditCard struct {
	number string
}

type Address struct{}
