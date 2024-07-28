package main

import (
	"fmt"
)

func main() {
	cashPayment := NewPayment("CASH", 100, "USD")
	cashPayment.Pay()
	fmt.Println()

	creditCardPayment := NewPayment("CREDIT CARD", 200, "USD")
	creditCardPayment.Pay()
	fmt.Println()

	cashPayment.Cancel()
	fmt.Println()

	creditCardPayment.Refund()
	fmt.Println()
}
