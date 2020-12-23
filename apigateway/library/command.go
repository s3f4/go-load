package library

import (
	"os/exec"

	"github.com/s3f4/go-load/apigateway/library/log"
)

// Command to run shell commands
type Command interface {
	Run(string) ([]byte, error)
}

type command struct {
}

// NewCommand returns command pointer
func NewCommand() Command {
	return &command{}
}

// Run runs multiple commands
func (*command) Run(command string) ([]byte, error) {
	cmd := exec.Command("/bin/sh", "-c", command)
	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Info(string(output))
		log.Info(err)
		return output, err
	}

	return output, nil
}
