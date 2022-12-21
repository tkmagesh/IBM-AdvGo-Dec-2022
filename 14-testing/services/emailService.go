package services

import "fmt"

type EmailService struct {
}

func (sms *EmailService) Send(msg string) bool {
	fmt.Printf("msg - [%q] sent as an email\n", msg)
	return true
}
