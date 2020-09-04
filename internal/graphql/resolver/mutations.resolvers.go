package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/vladimir-kopaliani/message-example/internal/models"
)

func (r *mutationResolver) WriteMessage(ctx context.Context, input models.WriteMessageInput) (*models.Message, error) {
	message, err := r.Service.WriteMessage(ctx, input)
	if err != nil {
		return nil, err
	}

	r.BroadcastChannel <- message

	return message, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
