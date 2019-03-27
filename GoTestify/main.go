package main

import (
	"fmt"
)

// Calculate returns x + 2.
func Calculate(x int) (result int) {
	result = x + 2
	return result
}

// MessageService handles notifying clients they have been charged
type MessageService interface {
	SendChargeNotification(int) bool
}

// SMSService is implementation of MessageService
type SMSService struct{}

// MyService uses the MessageService to notify clients
type MyService struct {
	messageService MessageService
}

// SendChargeNotification notifies clients they have been charged via SMS
// This is the method going to mock
func (sms SMSService) SendChargeNotification(value int) bool {
	fmt.Println("Sending Production Charge Notification")
	return true
}

// ChargeCustomer performs the charge to the customer
// In a real system we would maybe mock this as well
// but here, I want to make some money every time I run my tests
func (a MyService) ChargeCustomer(value int) error {
	a.messageService.SendChargeNotification(value)
	fmt.Printf("Charging Customer For the value of %d\n", value)
	return nil
}

func main() {
	fmt.Println("Hello World")

	smsService := SMSService{}
	myService := MyService{smsService}
	myService.ChargeCustomer(100)
}
