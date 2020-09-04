package resolver

import "github.com/vladimir-kopaliani/message-example/internal/models"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Service          Service
	BroadcastChannel chan *models.Message
}
