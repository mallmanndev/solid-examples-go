package main

import "fmt"

type PaymentGateway interface {
	Pay(p *Payment) error
}

type CancelablePaymentGateway interface {
	PaymentGateway
	Cancel(p *Payment) error
}

type RefundablePaymentGateway interface {
	PaymentGateway
	Refund(p *Payment) error
}

type CashPaymentGateway struct{}

func NewCashPaymentGateway() *CashPaymentGateway {
	return &CashPaymentGateway{}
}

func (pg *CashPaymentGateway) Pay(p *Payment) error {
	fmt.Printf("Paying %.2f %s with CASH\n", p.amount, p.currency)
	return nil
}

type PixPaymentGateway struct{}

func NewPixPaymentGateway() *PixPaymentGateway {
	return &PixPaymentGateway{}
}

func (pg *PixPaymentGateway) Pay(p *Payment) error {
	fmt.Printf("Paying with %.2f %s with PIX\n", p.amount, p.currency)
	return nil
}

func (pg *PixPaymentGateway) Refund(p *Payment) error {
	fmt.Printf("Refund PIX payment...\n")
	return nil
}

type TedPaymentGateway struct{}

func NewTedPaymentGateway() *TedPaymentGateway {
	return &TedPaymentGateway{}
}

func (pg *TedPaymentGateway) Pay(p *Payment) error {
	fmt.Printf("Paying with %.2f %s with TED\n", p.amount, p.currency)
	return nil
}

func (pg *TedPaymentGateway) Cancel(p *Payment) error {
	fmt.Printf("Cancel TED payment...\n")
	return nil
}
