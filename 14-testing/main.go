package main

import "testing-demo/services"

func main() {
	/*
		smsService := &services.SMSService{}
		mp := services.NewMessageProcessor(smsService)
	*/
	emailService := &services.EmailService{}
	mp := services.NewMessageProcessor(emailService)
	mp.Process("Parcel Received")
}
