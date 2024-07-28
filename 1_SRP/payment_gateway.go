package main

import "fmt"

type PaymentGateway struct{}

func NewPaymentGateway() *PaymentGateway {
	return &PaymentGateway{}
}

func (pg *PaymentGateway) Pay(p *Payment) error {
	switch p.method {
	case "CASH":
		fmt.Printf("Paying with %.2f %s with %s\n", p.amount, p.currency, p.method)
	case "PIX":
		fmt.Printf("Paying with %.2f %s with %s\n", p.amount, p.currency, p.method)
	case "TED":
		fmt.Printf("Paying with %.2f %s with %s\n", p.amount, p.currency, p.method)
	default:
		return fmt.Errorf("method %s not found", p.method)
	}
	return nil
}

func (pg *PaymentGateway) Cancel(p *Payment) error {
	fmt.Printf("Cancel %s payment...\n", p.method)
	return nil
}

func (pg *PaymentGateway) Refund(p *Payment) error {
	fmt.Printf("Refund %s payment...\n", p.method)
	return nil
}
