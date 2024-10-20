package chat

import "github.com/devenjarvis/pr-staton/internal/model"

type Api interface {
	SendMessage(*model.Message) error
}

func NewService(chatApi Api) *service {
	return &service{chatApi: chatApi}
}

type service struct {
	chatApi Api
}

func (s service) SendMessage(message *model.Message) error {
	return s.chatApi.SendMessage(message)
}
