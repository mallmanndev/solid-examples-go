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
	uc := NewPayUseCase(NewPaymentRepository(), NewNotificationsService(), NewPaymentGateway())

	err := uc.Execute("CASH", 100.0, "USD")

	assert.NoError(t, err)
}

func TestPayment_Pay_Pix(t *testing.T) {
	uc := NewPayUseCase(NewPaymentRepository(), NewNotificationsService(), NewPaymentGateway())

	err := uc.Execute("PIX", 200.0, "BRL")

	assert.NoError(t, err)
}

func TestPayment_Pay_Ted(t *testing.T) {
	uc := NewPayUseCase(NewPaymentRepository(), NewNotificationsService(), NewPaymentGateway())

	err := uc.Execute("TED", 300.0, "BRL")

	assert.NoError(t, err)
}

func TestPayment_Pay_InvalidMethod(t *testing.T) {
	uc := NewPayUseCase(NewPaymentRepository(), NewNotificationsService(), NewPaymentGateway())

	err := uc.Execute("INVALID", 400.0, "USD")

	assert.Error(t, err)
	assert.Equal(t, "method INVALID not found", err.Error())
}
