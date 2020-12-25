package app

import (
	"encoding/json"

	"github.com/s3f4/go-load/eventhandler/library/specs"
	"github.com/s3f4/go-load/eventhandler/models"
	"github.com/s3f4/go-load/eventhandler/repository"
	"github.com/s3f4/mu/log"
	"github.com/streadway/amqp"
)

// GetSpecs return queue specs
func GetSpecs(r repository.ResponseRepository) []*specs.ListenSpec {
	spec := &specs.ListenSpec{
		Consumer: "eventhandler",
		Queue:    "eventhandler",
		ProcessFunc: func(d *amqp.Delivery, exit chan<- struct{}) error {
			log.Infof("Received a message: %s", d.Body)
			var resp models.Response
			if err := json.Unmarshal(d.Body, &resp); err != nil {
				log.Error(err)
			}
			r.Create(&resp)
			return nil
		},
	}

	return []*specs.ListenSpec{
		spec,
	}
}
