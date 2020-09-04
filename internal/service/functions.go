package service

import (
	"context"
	er "github.com/vladimir-kopaliani/message-example/internal/internal_errors"
	"github.com/vladimir-kopaliani/message-example/internal/models"
	"time"
)

// WriteMessage ...
func (s *Service) WriteMessage(ctx context.Context, input models.WriteMessageInput) (*models.Message, error) {
	// TODO: validate input

	message := models.Message{
		ID:        input.ID,
		Text:      input.Text,
		CreatedAt: models.DateTime(time.Now().UTC()),
	}

	err := s.messagesRepository.SaveMessage(ctx, &message)
	if err != nil {
		return nil, err
	}

	return &message, nil
}

// GetMessages ...
func (s *Service) GetMessages(ctx context.Context, input models.GetMessagesInput) ([]*models.Message, error) {
	// validate input
	if time.Time(input.StartDate).After(time.Time(input.EndDate)) {
		return nil, er.ErrWrongInput
	}

	messages, err := s.messagesRepository.GetMessages(ctx, input)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
