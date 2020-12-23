package services

import (
	"errors"
	"fmt"
	"testing"

	"github.com/s3f4/go-load/apigateway/mocks"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Test_Start(t *testing.T) {
	ir := new(mocks.InstanceRepository)
	tr := new(mocks.TestRepository)
	rtr := new(mocks.RunTestRepository)
	q := new(mocks.QueueService)

	test := &models.Test{ID: 1, RequestCount: 10}
	instances := []models.InstanceTerraform{
		{ID: fmt.Sprintf("%d", 1)}, {ID: fmt.Sprintf("%d", 2)},
	}

	ir.On("GetFromTerraform").Return(instances, nil)
	rtr.On("Create", mock.Anything).Return(nil)
	rtr.On("Update", mock.Anything).Return(nil)
	q.On("Send", "worker", mock.Anything).Return(nil)
	q.On("Declare", mock.Anything).Return(nil)
	q.On("Delete", mock.Anything).Return(nil)
	q.On("Listen", mock.Anything).Return(nil)

	testService := NewTestService(ir, tr, rtr, q)
	runTestResult, err := testService.Start(test)

	assert.Nil(t, err)
	assert.NotNil(t, runTestResult)
}

func Test_Test_Start_Req(t *testing.T) {
	ir := new(mocks.InstanceRepository)
	tr := new(mocks.TestRepository)
	rtr := new(mocks.RunTestRepository)
	q := new(mocks.QueueService)

	test := &models.Test{ID: 1, RequestCount: 1}
	instances := []models.InstanceTerraform{
		{ID: fmt.Sprintf("%d", 1)}, {ID: fmt.Sprintf("%d", 2)},
	}

	ir.On("GetFromTerraform").Return(instances, nil)
	rtr.On("Create", mock.Anything).Return(nil)
	rtr.On("Update", mock.Anything).Return(nil)
	q.On("Send", "worker", mock.Anything).Return(nil)
	q.On("Declare", mock.Anything).Return(nil)
	q.On("Delete", mock.Anything).Return(nil)
	q.On("Listen", mock.Anything).Return(nil)

	testService := NewTestService(ir, tr, rtr, q)
	runTestResult, err := testService.Start(test)

	assert.Nil(t, err)
	assert.NotNil(t, runTestResult)
}

func Test_Test_Start_Req_SendMeessage_Error(t *testing.T) {
	ir := new(mocks.InstanceRepository)
	tr := new(mocks.TestRepository)
	rtr := new(mocks.RunTestRepository)
	q := new(mocks.QueueService)

	test := &models.Test{ID: 1, RequestCount: 1}
	instances := []models.InstanceTerraform{
		{ID: fmt.Sprintf("%d", 1)}, {ID: fmt.Sprintf("%d", 2)},
	}

	ir.On("GetFromTerraform").Return(instances, nil)
	rtr.On("Create", mock.Anything).Return(nil)
	rtr.On("Update", mock.Anything).Return(nil)
	q.On("Send", "worker", mock.Anything).Return(errors.New(""))
	q.On("Declare", mock.Anything).Return(nil)
	q.On("Delete", mock.Anything).Return(nil)
	q.On("Listen", mock.Anything).Return(nil)

	testService := NewTestService(ir, tr, rtr, q)
	runTestResult, err := testService.Start(test)

	assert.NotNil(t, err)
	assert.Nil(t, runTestResult)
}

func Test_Test_Start_TerraformError(t *testing.T) {
	ir := new(mocks.InstanceRepository)
	tr := new(mocks.TestRepository)
	rtr := new(mocks.RunTestRepository)
	q := new(mocks.QueueService)

	test := &models.Test{ID: 1, RequestCount: 10}

	ir.On("GetFromTerraform").Return(nil, errors.New(""))
	testService := NewTestService(ir, tr, rtr, q)
	runTestResult, err := testService.Start(test)

	assert.NotNil(t, err)
	assert.Nil(t, runTestResult)
}

func Test_Test_Start_RunTestCreateError(t *testing.T) {
	ir := new(mocks.InstanceRepository)
	tr := new(mocks.TestRepository)
	rtr := new(mocks.RunTestRepository)
	q := new(mocks.QueueService)

	test := &models.Test{ID: 1, RequestCount: 10}
	instances := []models.InstanceTerraform{
		{ID: fmt.Sprintf("%d", 1)}, {ID: fmt.Sprintf("%d", 2)},
	}

	ir.On("GetFromTerraform").Return(instances, nil)
	rtr.On("Create", mock.Anything).Return(errors.New(""))
	rtr.On("Update", mock.Anything).Return(nil)
	q.On("Send", "worker", mock.Anything).Return(nil)
	q.On("Declare", mock.Anything).Return(nil)
	q.On("Delete", mock.Anything).Return(nil)
	q.On("Listen", mock.Anything).Return(nil)

	testService := NewTestService(ir, tr, rtr, q)
	runTestResult, err := testService.Start(test)

	assert.NotNil(t, err)
	assert.Nil(t, runTestResult)
}

func Test_Test_SendError(t *testing.T) {
	ir := new(mocks.InstanceRepository)
	tr := new(mocks.TestRepository)
	rtr := new(mocks.RunTestRepository)
	q := new(mocks.QueueService)

	test := &models.Test{ID: 1, RequestCount: 10}
	instances := []models.InstanceTerraform{
		{ID: fmt.Sprintf("%d", 1)}, {ID: fmt.Sprintf("%d", 2)},
	}

	ir.On("GetFromTerraform").Return(instances, nil)
	rtr.On("Create", mock.Anything).Return(nil)
	rtr.On("Update", mock.Anything).Return(nil)
	q.On("Send", "worker", mock.Anything).Return(errors.New(""))
	q.On("Declare", mock.Anything).Return(nil)
	q.On("Delete", mock.Anything).Return(nil)
	q.On("Listen", mock.Anything).Return(nil)

	testService := NewTestService(ir, tr, rtr, q)
	runTestResult, err := testService.Start(test)

	assert.NotNil(t, err)
	assert.Nil(t, runTestResult)
}
