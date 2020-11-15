package library

import (
	"os/exec"

	"github.com/s3f4/go-load/apigateway/library/log"
)

// RunCommands runs multiple commands
func RunCommands(command string) ([]byte, error) {
	cmd := exec.Command("/bin/sh", "-c", command)
	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Info(string(output))
		log.Info(err)
		return output, err
	}
	
	return output, nil
}
