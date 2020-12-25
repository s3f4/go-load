package main

import (
	"github.com/s3f4/go-load/worker/app"
	"github.com/s3f4/go-load/worker/services"
)

func main() {
	queue := services.NewQueueService()
	worker := services.NewWorkerService()
	specs := app.GetSpecs(worker, queue)
	listener := services.NewListener()
	listener.Start(specs)
}
