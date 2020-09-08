package services

import (
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/s3f4/go-load/apigateway/template"
)

// InstanceServiceInterface ...
type InstanceServiceInterface interface {
	BuildTemplate(iReq models.Instance) error
	SpinUp() error
	Run() error
	Destroy() error
}

type instanceService struct {
}

var (
	// InstanceService handles all operations of instance handler
	InstanceService InstanceServiceInterface = &instanceService{}
)

func (is *instanceService) BuildTemplate(iReq models.Instance) error {
	t := template.NewInfraBuilder(
		iReq.Region,
		iReq.InstanceSize,
		iReq.InstanceCount,
	)

	if err := t.Write(); err != nil {
		return err
	}
	return nil
}
func (is *instanceService) SpinUp() error {
	return nil
}
func (is *instanceService) Run() error {
	return nil
}
func (is *instanceService) Destroy() error { return nil }
