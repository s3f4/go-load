package main

import "github.com/s3f4/go-load/worker/services"

func main() {
	listener := services.NewListener()
	queues := []string{"worker"}
	listener.Start(queues...)
}
