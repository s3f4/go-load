package repository

import "github.com/s3f4/go-load/eventhandler/models"

type ResponseRepository interface {
	Insert(*models.Response)
	Delete(*models.Response)
	List(query interface{}) []*models.Response
}