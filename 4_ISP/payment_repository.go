package main

import "fmt"

type PaymentRepository struct{}

func NewPaymentRepository() *PaymentRepository {
	return &PaymentRepository{}
}

func (pr *PaymentRepository) Save(p *Payment) error {
	fmt.Printf("Saving %s payment in database...\n", p.method)
	return nil
}
