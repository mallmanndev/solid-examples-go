package main

type CancelUseCase struct {
	paymentRepository    *PaymentRepository
	notificationsService *NotificationsService
	gateway              CancelablePaymentGateway
}

func NewCancelUseCase(pr *PaymentRepository, ns *NotificationsService, pg CancelablePaymentGateway) *CancelUseCase {
	return &CancelUseCase{
		paymentRepository:    pr,
		notificationsService: ns,
		gateway:              pg,
	}
}

func (p *CancelUseCase) Execute(method string, value float64, currency string) error {
	payment := NewPayment(method, value, currency)
	err := p.gateway.Cancel(payment)
	if err != nil {
		return err
	}

	err = p.paymentRepository.Save(payment)
	if err != nil {
		return err
	}

	err = p.notificationsService.SendNotification("Refund executed")
	if err != nil {
		return err
	}

	return nil
}
