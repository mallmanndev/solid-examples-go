package main

import "errors"

type PaymentWithValidationUseCase struct {
	paymentUseCase *PayUseCase
}

func NewPaymentWithValidationUseCase(pr *PaymentRepository, ns *NotificationsService, pg PaymentGateway) *PaymentWithValidationUseCase {
	puc := NewPayUseCase(pr, ns, pg)
	return &PaymentWithValidationUseCase{
		paymentUseCase: puc,
	}
}

func (p *PaymentWithValidationUseCase) Execute(method string, value float64, currency string) error {
	// Validate payment
	if value <= 0 {
		return errors.New("invalid payment value")
	}

	return p.paymentUseCase.Execute(method, value, currency)
}
