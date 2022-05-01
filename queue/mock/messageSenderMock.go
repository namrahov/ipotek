package mock

import "github.com/stretchr/testify/mock"

type MessageSenderMock struct {
	mock.Mock
}

func (m *MessageSenderMock) SendMessage(message interface{}, queueName string) error {
	args := m.Called(message, queueName)
	return args.Error(0)
}
