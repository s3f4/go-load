package services

import (
	"testing"

	"github.com/s3f4/go-load/apigateway/mocks"
	"github.com/s3f4/go-load/apigateway/models"
)

func Test_Instance_Show(t *testing.T) {
	command := new(mocks.Command)
	command.On("Run", "cd infra;terraform output -json account").Return([]byte("abc"), nil)
	ir := new(mocks.InstanceRepository)
	instanceService := NewInstanceService(ir, command)
	instanceService.Show("account")
}

func Test_Instance_BuildTemplate(t *testing.T) {
	ir := new(mocks.InstanceRepository)
	command := new(mocks.Command)
	is := NewInstanceService(ir, command)
	is.BuildTemplate(models.InstanceConfig{})

}
