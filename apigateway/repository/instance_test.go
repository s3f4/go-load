package repository

import (
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/s3f4/go-load/apigateway/library"
	"github.com/s3f4/go-load/apigateway/mocks"
	"github.com/s3f4/go-load/apigateway/models"
	"github.com/stretchr/testify/assert"
)

func Test_Instance_Create(t *testing.T) {
	_, sqlMock, conn := ConnectMock(MYSQL)
	r := NewInstanceRepository(conn, library.NewCommand())
	instance := &models.InstanceConfig{
		ID: 1,
	}

	sqlMock.ExpectExec(regexp.QuoteMeta("DELETE FROM `instance_configs` WHERE 1=1")).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectExec(regexp.QuoteMeta("DELETE FROM `instances` WHERE 1=1")).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `instance_configs` (`id`) VALUES (?)")).
		WithArgs(
			instance.ID,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := r.Create(instance)
	assert.Nil(t, err)
}

func Test_Instance_Get(t *testing.T) {
	_, sqlMock, conn := ConnectMock(MYSQL)
	r := NewInstanceRepository(conn, library.NewCommand())
	instance := &models.InstanceConfig{
		ID: 1,
		Configs: []*models.Instance{
			{
				ID:               1,
				InstanceConfigID: 1,
			},
		},
	}

	instanceRows := sqlmock.NewRows([]string{"id", "instance_config_id"}).
		AddRow(1, 1)

	instanceConfigs := sqlmock.NewRows([]string{"id"}).
		AddRow(1)

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `instance_configs` ORDER BY `instance_configs`.`id` DESC LIMIT 1")).
		WillReturnRows(instanceConfigs)

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `instances` WHERE `instances`.`instance_config_id` = ?")).
		WithArgs(instance.ID).
		WillReturnRows(instanceRows)

	instanceResult, err := r.Get()
	assert.Equal(t, instance, instanceResult)
	assert.Nil(t, err)
}

func Test_Instance_Get_Error(t *testing.T) {
	_, sqlMock, conn := ConnectMock(MYSQL)
	r := NewInstanceRepository(conn, library.NewCommand())

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `instance_configs` ORDER BY `instance_configs`.`id` DESC LIMIT 1")).
		WillReturnError(errors.New(""))

	_, err := r.Get()
	assert.NotNil(t, err)
}

func Test_Instance_Delete(t *testing.T) {
	_, sqlMock, conn := ConnectMock(MYSQL)
	r := NewInstanceRepository(conn, library.NewCommand())
	instance := &models.InstanceConfig{ID: 1}

	sqlMock.ExpectExec(regexp.QuoteMeta("DELETE FROM `instance_configs` WHERE 1=1 AND `instance_configs`.`id` = ?")).
		WithArgs(instance.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := r.Delete(instance)
	fmt.Println(err)
	assert.Nil(t, err)
}

func Test_Instance_GetFromTerraform(t *testing.T) {
	output := []byte(`{"0":{"backups":false,"created_at":null,"disk":null,"id":null,"image":"ubuntu-18-04-x64","ipv4_address":null,"ipv4_address_private":null,"ipv6":false,"ipv6_address":null,"locked":null,"memory":null,"monitoring":false,"name":"worker-nyc1-1","price_hourly":null,"price_monthly":null,"private_networking":null,"region":"nyc1","resize_disk":true,"size":"s-1vcpu-1gb","ssh_keys":["2"],"status":null,"tags":null,"urn":null,"user_data":null,"vcpus":null,"volume_ids":null,"vpc_uuid":null},"1":{"backups":false,"created_at":null,"disk":null,"id":null,"image":"ubuntu-18-04-x64","ipv4_address":null,"ipv4_address_private":null,"ipv6":false,"ipv6_address":null,"locked":null,"memory":null,"monitoring":false,"name":"worker-sgp1-1","price_hourly":null,"price_monthly":null,"private_networking":null,"region":"sgp1","resize_disk":true,"size":"s-1vcpu-1gb","ssh_keys":["2"],"status":null,"tags":null,"urn":null,"user_data":null,"vcpus":null,"volume_ids":null,"vpc_uuid":null}}`)
	instancesResult := []models.InstanceTerraform{
		{
			ID:                 "",
			Name:               "worker-nyc1-1",
			CreatedAt:          "",
			Disk:               0,
			Image:              "ubuntu-18-04-x64",
			IPV4Address:        "",
			IPV4AddressPrivate: "",
			Memory:             0,
			Region:             "nyc1",
			Size:               "s-1vcpu-1gb",
			Status:             "",
		},
		{
			ID:                 "",
			Name:               "worker-sgp1-1",
			CreatedAt:          "",
			Disk:               0,
			Image:              "ubuntu-18-04-x64",
			IPV4Address:        "",
			IPV4AddressPrivate: "",
			Memory:             0,
			Region:             "sgp1",
			Size:               "s-1vcpu-1gb",
			Status:             "",
		},
	}
	_, _, conn := ConnectMock(MYSQL)
	command := new(mocks.Command)
	r := NewInstanceRepository(conn, command)
	command.On("Run", "cd infra;terraform output -json workers").Return(output, nil)
	instances, _ := r.GetFromTerraform()
	assert.Equal(t, instancesResult, instances)
}

func Test_Instance_GetFromTerraform_CommandError(t *testing.T) {
	_, _, conn := ConnectMock(MYSQL)
	command := new(mocks.Command)
	r := NewInstanceRepository(conn, command)
	command.On("Run", "cd infra;terraform output -json workers").Return(nil, errors.New(""))
	_, err := r.GetFromTerraform()
	assert.NotNil(t, err)
}

func Test_Instance_GetFromTerraform_JSONError(t *testing.T) {
	output := []byte(`{"0":{"backups":false,"created_at":null,"disk":null,"id":null,"image":"ubuntu-18-04-x64","ipv4_address":null,"ipv4_address_private":null,"ipv6":false,"ipv6_address":null,"locked":null,"memory":null,"monitoring":false,"name":"worker-nyc1-1","price_hourly":null,"price_monthly":null,"private_networking":null,"region":"nyc1","resize_disk":true,"size":"s-1vcpu-1gb","ssh_keys":["2"],"status":null,"tags":null,"urn":null,"user_data":null,"vcpus":null,"volume_ids":null,"vpc_uuid":null},"1":{"backups":false,"created_at":null,"disk":null,"id":null,"image":"ubuntu-18-04-x64","ipv4_address":null,"ipv4_address_private":null,"ipv6":false,"ipv6_address":null,"locked":null,"memory":null,"monitoring":false,"name":"worker-sgp1-1","price_hourly":null,"price_monthly":null,"private_networking":null,"region":"sgp1","resize_disk":true,"size":"s-1vcpu-1gb","ssh_keys":["2"],"status":null,"tags":null,"urn":null,"user_data":null,"vcpus":null,"volume_ids":null,"vpc_uuid":null}`)
	_, _, conn := ConnectMock(MYSQL)
	command := new(mocks.Command)
	r := NewInstanceRepository(conn, command)
	command.On("Run", "cd infra;terraform output -json workers").Return(output, nil)
	_, err := r.GetFromTerraform()
	assert.NotNil(t, err)
}

func Test_Instance_GetFromTerraform_ZeroResult(t *testing.T) {
	output := []byte(`{}`)
	_, _, conn := ConnectMock(MYSQL)
	command := new(mocks.Command)
	r := NewInstanceRepository(conn, command)
	command.On("Run", "cd infra;terraform output -json workers").Return(output, nil)
	_, err := r.GetFromTerraform()
	assert.NotNil(t, err)
}
