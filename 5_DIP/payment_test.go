package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPayment(t *testing.T) {
	p := NewPayment("CASH", 100.0, "USD")

	assert.Equal(t, "CASH", p.method)
	assert.Equal(t, 100.0, p.amount)
	assert.Equal(t, "USD", p.currency)
}

func TestPayment_Pay_Cash(t *testing.T) {
	uc := NewPayUseCase(NewPaymentRepository(), NewNotificationsService(), NewCashPaymentGateway())

	err := uc.Execute("CASH", 100.0, "USD")

	assert.NoError(t, err)
}

func TestPayment_Pay_Pix(t *testing.T) {
	uc := NewPayUseCase(NewPaymentRepository(), NewNotificationsService(), NewPixPaymentGateway())

	err := uc.Execute("PIX", 200.0, "BRL")

	assert.NoError(t, err)
}

func TestPayment_Pay_Ted(t *testing.T) {
	uc := NewPayUseCase(NewPaymentRepository(), NewNotificationsService(), NewTedPaymentGateway())

	err := uc.Execute("TED", 300.0, "BRL")

	assert.NoError(t, err)
}
