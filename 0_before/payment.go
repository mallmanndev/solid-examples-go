package main

import "fmt"

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

func (p *Payment) Pay() error {
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
	if err := p.save(); err != nil {
		return err
	}
	if err := p.sendNotification(); err != nil {
		return err
	}
	return nil
}

func (p *Payment) Cancel() error {
	fmt.Printf("Cancel %s payment...\n", p.method)
	if err := p.save(); err != nil {
		return err
	}
	if err := p.sendNotification(); err != nil {
		return err
	}
	return nil
}

func (p *Payment) Refund() error {
	fmt.Printf("Refund %s payment...\n", p.method)
	if err := p.save(); err != nil {
		return err
	}
	if err := p.sendNotification(); err != nil {
		return err
	}
	return nil
}

func (p *Payment) save() error {
	fmt.Printf("Saving %s payment in database...\n", p.method)
	return nil
}

func (p *Payment) sendNotification() error {
	fmt.Println("Send notification...")
	return nil
}
