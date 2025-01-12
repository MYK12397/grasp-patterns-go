package main

import "fmt"

type PaymentService interface {
	ProcessPayment(amount float64) string
}

type PaypalService struct{}

func (p *PaypalService) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Processed payment of $%.2f through paypal", amount)
}

type StripeService struct{}

func (s *StripeService) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Processed payment of $%.2f through stripe", amount)
}

type PaymentProcessor struct {
	paymentService PaymentService
}

func NewPaymentProcessor(service PaymentService) *PaymentProcessor {
	return &PaymentProcessor{paymentService: service}
}

func (pp *PaymentProcessor) Pay(amount float64) {
	result := pp.paymentService.ProcessPayment(amount)
	fmt.Println(result)
}

func main() {
	paypal := &PaypalService{}
	paymentProcessor := NewPaymentProcessor(paypal)
	paymentProcessor.Pay(120.0)

	stripe := &StripeService{}
	paymentProcessor = NewPaymentProcessor(stripe)

	paymentProcessor.Pay(200.0)

}
