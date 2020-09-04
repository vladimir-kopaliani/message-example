//go:generate go run github.com/99designs/gqlgen

package resolver

import (
	"context"
	"github.com/vladimir-kopaliani/message-example/internal/models"
)

// Service define functions which service has to implement
type Service interface {
	WriteMessage(ctx context.Context, input models.WriteMessageInput) (*models.Message, error)
	GetMessages(ctx context.Context, input models.GetMessagesInput) ([]*models.Message, error)
}
