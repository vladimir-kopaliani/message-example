package service

// Service represent service
type Service struct {
	logger             Logger
	messagesRepository MessagesRepository
}

// Configuration contains settings for service
type Configuration struct {
	Logger             Logger
	MessagesRepository MessagesRepository
}

// NewService creates new instance of service
func NewService(conf *Configuration) *Service {
	if conf.MessagesRepository == nil {
		panic("UsersRepository is nil")
	}
	if conf.Logger == nil {
		panic("Logger is not set")
	}

	return &Service{
		messagesRepository: conf.MessagesRepository,
		logger:             conf.Logger,
	}
}
