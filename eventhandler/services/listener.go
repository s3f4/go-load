package services

import "gorm.io/gorm"

// ListenerService is used to listen all queues
type ListenerService interface {
	Start(queues ...string)
}

// Listener listens all queues
type listener struct {
	service *rabbitMQService
}

// NewListener returns new listener
func NewListener(db *gorm.DB) ListenerService {
	return &listener{service: NewRabbitMQService(db).(*rabbitMQService)}
}

// Start starts listening all queues.
func (l *listener) Start(queues ...string) {
	block := make(chan struct{})
	for _, queue := range queues {
		go l.service.Listen(queue)
	}
	<-block
}
