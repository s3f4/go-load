package specs

import "github.com/streadway/amqp"

type processFunc func(d *amqp.Delivery, exit chan<- struct{}) error

// ListenSpec ...
type ListenSpec struct {
	Consumer    string
	Queue       string
	ProcessFunc processFunc
}
