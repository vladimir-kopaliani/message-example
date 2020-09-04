package service

import (
	"context"
	"github.com/vladimir-kopaliani/message-example/internal/models"
)

// Logger ...
type Logger interface {
	Debug(msg ...interface{})
	Info(msg ...interface{})
	Error(msg ...interface{})
	Fatal(msg ...interface{})
}

// MessagesRepository represent a message repository
type MessagesRepository interface {
	SaveMessage(ctx context.Context, message *models.Message) error
	GetMessages(ctx context.Context, input models.GetMessagesInput) ([]*models.Message, error)
}
