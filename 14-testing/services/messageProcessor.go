package services

type MessageProcessor struct {
	msgService MessageService
}

func (mp *MessageProcessor) Process(msg string) {
	mp.msgService.Send(msg)
	// mp.msgService.Send("dummy message")
}

func NewMessageProcessor(mService MessageService) *MessageProcessor {
	return &MessageProcessor{
		msgService: mService,
	}
}
