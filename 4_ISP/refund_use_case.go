package main

type RefundUseCase struct {
	paymentRepository    *PaymentRepository
	notificationsService *NotificationsService
	gateway              RefundablePaymentGateway
}

func NewRefundUseCase(pr *PaymentRepository, ns *NotificationsService, pg RefundablePaymentGateway) *RefundUseCase {
	return &RefundUseCase{
		paymentRepository:    pr,
		notificationsService: ns,
		gateway:              pg,
	}
}

func (p *RefundUseCase) Execute(method string, value float64, currency string) error {
	payment := NewPayment(method, value, currency)
	err := p.gateway.Refund(payment)
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
