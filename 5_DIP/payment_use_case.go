package main

type PayUseCase struct {
	paymentRepository    PaymentRepository
	notificationsService NotificationsService
	paymentGateway       PaymentGateway
}

func NewPayUseCase(pr PaymentRepository, ns NotificationsService, pg PaymentGateway) *PayUseCase {
	return &PayUseCase{
		paymentRepository:    pr,
		notificationsService: ns,
		paymentGateway:       pg,
	}
}

func (p *PayUseCase) Execute(method string, value float64, currency string) error {
	payment := NewPayment(method, value, currency)
	err := p.paymentGateway.Pay(payment)
	if err != nil {
		return err
	}

	err = p.paymentRepository.Save(payment)
	if err != nil {
		return err
	}

	err = p.notificationsService.SendNotification("Payment executed")
	if err != nil {
		return err
	}

	return nil
}
