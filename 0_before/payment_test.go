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
	p := NewPayment("CASH", 100.0, "USD")

	err := p.Pay()

	assert.NoError(t, err)
}

func TestPayment_Pay_Pix(t *testing.T) {
	p := NewPayment("PIX", 200.0, "BRL")

	err := p.Pay()

	assert.NoError(t, err)
}

func TestPayment_Pay_Ted(t *testing.T) {
	p := NewPayment("TED", 300.0, "BRL")

	err := p.Pay()

	assert.NoError(t, err)
}

func TestPayment_Pay_InvalidMethod(t *testing.T) {
	p := NewPayment("INVALID", 400.0, "USD")

	err := p.Pay()

	assert.Error(t, err)
	assert.Equal(t, "method INVALID not found", err.Error())
}

func TestPayment_Cancel(t *testing.T) {
	p := NewPayment("PIX", 100.0, "BRL")

	err := p.Cancel()

	assert.NoError(t, err)
}

func TestPayment_Refund(t *testing.T) {
	p := NewPayment("TED", 150.0, "BRL")

	err := p.Refund()

	assert.NoError(t, err)
}
