package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/vladimir-kopaliani/message-example/internal/models"
)

func (r *subscriptionResolver) NewMessagesSubscription(ctx context.Context) (<-chan *models.Message, error) {
	return r.BroadcastChannel, nil
}

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
