package main

import (
	"fmt"
	"go.uber.org/dig"
)

func main() {
	container := dig.New()
	err := container.Provide(NewMessageRepository)
	if err != nil {
		panic(err)
	}
	err = container.Provide(NewMessageService)
	if err != nil {
		panic(err)
	}

	var messageService *MessageService
	err = container.Invoke(func(service *MessageService) {
		messageService = service
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(messageService.GetMessage("World")) // Output: Hello, World!
}

type MessageService struct {
	messageRepository *MessageRepository
}

func NewMessageService(messageRepository *MessageRepository) *MessageService {
	return &MessageService{
		messageRepository: messageRepository,
	}
}

func (r *MessageService) GetMessage(message string) string {
	return r.messageRepository.GetMessage(message)
}

type MessageRepository struct {
}

func NewMessageRepository() *MessageRepository {
	return &MessageRepository{}
}

func (r *MessageRepository) GetMessage(message string) string {
	return fmt.Sprintf("Hello, %s!", message)
}
