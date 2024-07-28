package main

import "fmt"

type PaymentRepository interface {
	Save(p *Payment) error
}

type PaymentRepositoryImpl struct{}

func NewPaymentRepository() *PaymentRepositoryImpl {
	return &PaymentRepositoryImpl{}
}

func (pr *PaymentRepositoryImpl) Save(p *Payment) error {
	fmt.Printf("Saving %s payment in database...\n", p.method)
	return nil
}
