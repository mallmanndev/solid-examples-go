package main

import (
	"fmt"
)

func main() {
	uc := NewPayUseCase(NewPaymentRepository(), NewNotificationsService(), NewPaymentGateway())

	uc.Execute("CASH", 100, "USD")
	fmt.Println()

	uc.Execute("CASH", 100, "USD")
	fmt.Println()
}
