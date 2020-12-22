package middlewares

import (
	"github.com/s3f4/go-load/apigateway/repository"
	"github.com/s3f4/go-load/apigateway/services"
)

// Middleware holds dependencies of middlewares
type Middleware struct {
	tokenService        services.TokenService
	authService         services.AuthService
	testRepository      repository.TestRepository
	testGroupRepository repository.TestGroupRepository
	runTestRespository  repository.RunTestRepository
}

// NewMiddleware returns a new middleware object
func NewMiddleware(
	ts services.TokenService,
	as services.AuthService,
	tr repository.TestRepository,
	tgr repository.TestGroupRepository,
	rtr repository.RunTestRepository,
) *Middleware {
	return &Middleware{
		tokenService:        ts,
		authService:         as,
		testRepository:      tr,
		testGroupRepository: tgr,
		runTestRespository:  rtr,
	}
}
