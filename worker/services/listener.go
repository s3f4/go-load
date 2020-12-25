package services

import (
	"github.com/s3f4/go-load/worker/library/specs"
)

// ListenerService is used to listen all queues
type ListenerService interface {
	Start(specs []*specs.ListenSpec)
}

// Listener listens all queues
type listener struct {
	service *queueService
}

// NewListener returns new listener
func NewListener() ListenerService {
	return &listener{service: NewQueueService().(*queueService)}
}

// Start starts listening all queues.
func (l *listener) Start(specs []*specs.ListenSpec) {
	block := make(chan struct{})
	for _, spec := range specs {
		go l.service.Listen(spec)
	}
	<-block
}
