package main

type Payment struct {
	method   string
	amount   float64
	currency string
}

func NewPayment(method string, amount float64, currency string) *Payment {
	return &Payment{
		method:   method,
		amount:   amount,
		currency: currency,
	}
}
