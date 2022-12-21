package services

import "fmt"

type SMSService struct {
}

func (sms *SMSService) Send(msg string) bool {
	fmt.Printf("msg - [%q] sent as sms\n", msg)
	return true
}
