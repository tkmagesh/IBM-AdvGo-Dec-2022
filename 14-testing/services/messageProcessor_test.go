package services

import (
	"testing"
	mocks "testing-demo/mocks/services"
)

//custom Mock implementation
/*
type MockMessageService struct {
	sendInvoked bool
	msg         string
	returnValue bool
}

func (mms *MockMessageService) Send(msg string) bool {
	mms.sendInvoked = true
	mms.msg = msg
	return mms.returnValue
}

func Test_MessageProcessor(t *testing.T) {
	//arrange
	mockMessageService := &MockMessageService{
		returnValue: true,
	}
	sut := NewMessageProcessor(mockMessageService)
	testMsg := "Test Message"

	//act
	sut.Process(testMsg)

	//assert
	if !mockMessageService.sendInvoked {
		t.Error("Should send the message to the Message service")
	}
	if mockMessageService.msg != testMsg {
		t.Error("Wrong message sent to the Message service")
	}

}
*/
func Test_MessageProcesser_Sends_Message_To_MessageService(t *testing.T) {
	// arrange
	testMessage := "test message"
	mockMessageService := &mocks.MessageService{}

	//configure the mock
	mockMessageService.On("Send", testMessage).Return(true)
	sut := NewMessageProcessor(mockMessageService)

	// act
	sut.Process(testMessage)

	// assert
	mockMessageService.AssertExpectations(t)
}
