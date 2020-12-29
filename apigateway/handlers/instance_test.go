package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/docker/docker/api/types/swarm"
	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/mocks"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/stretchr/testify/assert"
)

func Test_Instance_SpinUp(t *testing.T) {
	service := new(mocks.InstanceService)
	instanceConfig := &models.InstanceConfig{ID: 1, Configs: []*models.Instance{
		{ID: 1},
	}}

	service.On("BuildTemplate", *instanceConfig).Return(4, nil)
	service.On("ScaleWorkers", 4).Return(nil)
	service.On("SpinUp").Return(nil)

	instanceHandler := NewInstanceHandler(service)

	res, body := makeRequest("/instances", http.MethodPost, instanceHandler.SpinUp, strings.NewReader(`{"id":1, "configs":[{"id":1}]}`))

	assert.Equal(t, `{"status":true,"data":{"id":1,"configs":[{"id":1,"instance_config_id":0,"count":0,"size":"","image":"","region":""}]}}`, string(body))
	assert.Equal(t, res.StatusCode, http.StatusOK, "%d status is not equal to %d", res.StatusCode, http.StatusOK)
}

func Test_Instance_SpinUp_ParseError(t *testing.T) {
	service := new(mocks.InstanceService)
	instanceConfig := &models.InstanceConfig{ID: 1, Configs: []*models.Instance{
		{ID: 1},
	}}

	service.On("BuildTemplate", *instanceConfig).Return(4, nil)
	service.On("ScaleWorkers", 4).Return(nil)
	service.On("SpinUp").Return(nil)

	instanceHandler := NewInstanceHandler(service)

	res, body := makeRequest("/instances", http.MethodPost, instanceHandler.SpinUp, strings.NewReader(`{"id":1, "configs":[{"id":1}]`))

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrBadRequest), string(body))
	assert.Equal(t, http.StatusBadRequest, res.StatusCode, "%d status is not equal to %d", http.StatusBadRequest, res.StatusCode)
}

func Test_Instance_SpinUp_BuildTemplateError(t *testing.T) {
	service := new(mocks.InstanceService)
	instanceConfig := &models.InstanceConfig{ID: 1, Configs: []*models.Instance{
		{ID: 1},
	}}

	service.On("BuildTemplate", *instanceConfig).Return(0, errors.New(library.ErrInternalServerError.Error()))
	service.On("ScaleWorkers", 4).Return(nil)
	service.On("SpinUp").Return(nil)

	instanceHandler := NewInstanceHandler(service)

	res, body := makeRequest("/instances", http.MethodPost, instanceHandler.SpinUp, strings.NewReader(`{"id":1, "configs":[{"id":1}]}`))

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrInternalServerError), string(body))
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "%d status is not equal to %d", http.StatusInternalServerError, res.StatusCode)
}

func Test_Instance_SpinUp_SpinUpError(t *testing.T) {
	service := new(mocks.InstanceService)
	instanceConfig := &models.InstanceConfig{ID: 1, Configs: []*models.Instance{
		{ID: 1},
	}}

	service.On("BuildTemplate", *instanceConfig).Return(4, nil)
	service.On("ScaleWorkers", 4).Return(nil)
	service.On("SpinUp").Return(errors.New(library.ErrInternalServerError.Error()))

	instanceHandler := NewInstanceHandler(service)

	res, body := makeRequest("/instances", http.MethodPost, instanceHandler.SpinUp, strings.NewReader(`{"id":1, "configs":[{"id":1}]}`))

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrInternalServerError), string(body))
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "%d status is not equal to %d", http.StatusInternalServerError, res.StatusCode)
}

func Test_Instance_SpinUp_ScaleWorkersError(t *testing.T) {
	service := new(mocks.InstanceService)
	instanceConfig := &models.InstanceConfig{ID: 1, Configs: []*models.Instance{
		{ID: 1},
	}}

	service.On("BuildTemplate", *instanceConfig).Return(4, nil)
	service.On("ScaleWorkers", 4).Return(errors.New(library.ErrInternalServerError.Error()))
	service.On("SpinUp").Return(nil)

	instanceHandler := NewInstanceHandler(service)
	os.Setenv("APP_ENV", "p")
	res, body := makeRequest("/instances", http.MethodPost, instanceHandler.SpinUp, strings.NewReader(`{"id":1, "configs":[{"id":1}]}`))

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrInternalServerError), string(body))
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "%d status is not equal to %d", http.StatusInternalServerError, res.StatusCode)
}

func Test_Instance_Destroy(t *testing.T) {
	service := new(mocks.InstanceService)
	service.On("Destroy").Return(nil)
	instanceHandler := NewInstanceHandler(service)

	res, body := makeRequest("/instances", http.MethodDelete, instanceHandler.Destroy, strings.NewReader(`{"id":1}`))

	assert.Equal(t, `{"status":true,"data":"Workers destroyed"}`, string(body))
	assert.Equal(t, res.StatusCode, http.StatusOK, "%d status is not equal to %d", res.StatusCode, http.StatusOK)
}

func Test_Instance_Destroy_Error(t *testing.T) {
	service := new(mocks.InstanceService)
	service.On("Destroy").Return(errors.New(library.ErrInternalServerError.Error()))
	instanceHandler := NewInstanceHandler(service)

	res, body := makeRequest("/instances", http.MethodDelete, instanceHandler.Destroy, strings.NewReader(`{"id":1}`))

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrInternalServerError), string(body))
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "%d status is not equal to %d", http.StatusInternalServerError, res.StatusCode)
}

func Test_Instance_ShowRegions(t *testing.T) {
	service := new(mocks.InstanceService)
	service.On("ShowRegions").Return("output", nil)
	instanceHandler := NewInstanceHandler(service)

	res, body := makeRequest("/instances/regions", http.MethodGet, instanceHandler.ShowRegions, nil)

	assert.Equal(t, `{"status":true,"data":"output"}`, string(body))
	assert.Equal(t, res.StatusCode, http.StatusOK, "%d status is not equal to %d", res.StatusCode, http.StatusOK)
}

func Test_Instance_ShowRegions_Error(t *testing.T) {
	service := new(mocks.InstanceService)
	service.On("ShowRegions").Return("", errors.New(""))
	instanceHandler := NewInstanceHandler(service)

	res, body := makeRequest("/instances/regions", http.MethodGet, instanceHandler.ShowRegions, nil)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrInternalServerError), string(body))
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "%d status is not equal to %d", http.StatusInternalServerError, res.StatusCode)
}

func Test_Instance_ShowAccount(t *testing.T) {
	service := new(mocks.InstanceService)
	service.On("ShowAccount").Return("output", nil)
	instanceHandler := NewInstanceHandler(service)

	res, body := makeRequest("/instances/account", http.MethodGet, instanceHandler.ShowAccount, nil)

	assert.Equal(t, `{"status":true,"data":"output"}`, string(body))
	assert.Equal(t, res.StatusCode, http.StatusOK, "%d status is not equal to %d", res.StatusCode, http.StatusOK)
}

func Test_Instance_ShowAccount_Error(t *testing.T) {
	service := new(mocks.InstanceService)
	service.On("ShowAccount").Return("", errors.New(""))
	instanceHandler := NewInstanceHandler(service)

	res, body := makeRequest("/instances/account", http.MethodGet, instanceHandler.ShowAccount, nil)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrInternalServerError), string(body))
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "%d status is not equal to %d", http.StatusInternalServerError, res.StatusCode)
}

func Test_Instance_ShowSwarmNodes(t *testing.T) {
	service := new(mocks.InstanceService)
	service.On("ShowSwarmNodes").Return([]swarm.Node{}, nil)
	instanceHandler := NewInstanceHandler(service)

	res, body := makeRequest("/instances/swarm-nodes", http.MethodGet, instanceHandler.ShowSwarmNodes, nil)

	assert.Equal(t, fmt.Sprintf(`{"status":true,"data":%v}`, []swarm.Node{}), string(body))
	assert.Equal(t, res.StatusCode, http.StatusOK, "%d status is not equal to %d", res.StatusCode, http.StatusOK)
}

func Test_Instance_ShowSwarmNodes_Error(t *testing.T) {
	service := new(mocks.InstanceService)
	service.On("ShowSwarmNodes").Return(nil, errors.New(""))
	instanceHandler := NewInstanceHandler(service)

	res, body := makeRequest("/instances/swarm-nodes", http.MethodGet, instanceHandler.ShowSwarmNodes, nil)

	assert.Equal(t, `{"status":false}`, string(body))
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "%d status is not equal to %d", http.StatusInternalServerError, res.StatusCode)
}

func Test_Instance_GetInstanceInfo(t *testing.T) {
	service := new(mocks.InstanceService)
	service.On("GetInstanceInfo").Return(&models.InstanceConfig{}, nil)
	instanceHandler := NewInstanceHandler(service)

	res, body := makeRequest("/instances", http.MethodGet, instanceHandler.GetInstanceInfo, nil)
	instanceConfigStr, _ := json.Marshal(&models.InstanceConfig{})
	assert.Equal(t, fmt.Sprintf(`{"status":true,"data":%s}`, instanceConfigStr), string(body))
	assert.Equal(t, res.StatusCode, http.StatusOK, "%d status is not equal to %d", res.StatusCode, http.StatusOK)
}

func Test_Instance_GetInstanceInfo_Error(t *testing.T) {
	service := new(mocks.InstanceService)
	service.On("GetInstanceInfo").Return(nil, errors.New(""))
	instanceHandler := NewInstanceHandler(service)

	res, body := makeRequest("/instances", http.MethodGet, instanceHandler.GetInstanceInfo, nil)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrNotFound), string(body))
	assert.Equal(t, http.StatusNotFound, res.StatusCode, "%d status is not equal to %d", http.StatusNotFound, res.StatusCode)
}

func Test_Instance_GetInstanceInfoFromTerraform(t *testing.T) {
	service := new(mocks.InstanceService)
	service.On("GetInstanceInfoFromTerraform").Return("output", nil)
	instanceHandler := NewInstanceHandler(service)

	res, body := makeRequest("/instances/terraform", http.MethodGet, instanceHandler.GetInstanceInfoFromTerraform, nil)

	assert.Equal(t, `{"status":true,"data":"output"}`, string(body))
	assert.Equal(t, res.StatusCode, http.StatusOK, "%d status is not equal to %d", res.StatusCode, http.StatusOK)
}

func Test_Instance_GetInstanceInfoFromTerraform_Error(t *testing.T) {
	service := new(mocks.InstanceService)
	service.On("GetInstanceInfoFromTerraform").Return("", errors.New(""))
	instanceHandler := NewInstanceHandler(service)

	res, body := makeRequest("/instances/terraform", http.MethodGet, instanceHandler.GetInstanceInfoFromTerraform, nil)

	assert.Equal(t, fmt.Sprintf(`{"status":false,"message":"%s"}`, library.ErrNotFound), string(body))
	assert.Equal(t, http.StatusNotFound, res.StatusCode, "%d status is not equal to %d", http.StatusNotFound, res.StatusCode)
}
