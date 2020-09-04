package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/vladimir-kopaliani/message-example/internal/models"
)

func (r *queryResolver) GetMessages(ctx context.Context, input models.GetMessagesInput) ([]*models.Message, error) {
	return r.Service.GetMessages(ctx, input)
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
