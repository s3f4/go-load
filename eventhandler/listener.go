package eventhandler

// ListenerService is used to listen all queues
type ListenerService interface {
	Start(queues ...string)
}

// Listener listens all queues
type listener struct {
	service *rabbitMQService
}

// NewListener returns new listener
func NewListener(uri string) ListenerService {
	return &listener{service: NewRabbitMQService(uri).(*rabbitMQService)}
}

// Start starts listening all queues.
func (l *listener) Start(queues ...string) {
	for _, queue := range queues {
		l.service.Listen(queue)
	}
}
