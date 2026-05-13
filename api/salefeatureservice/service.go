package salefeatureservice

import "github.com/gofurry/steam-go/internal/request"

// Service exposes ISaleFeatureService methods.
type Service struct {
	executor *request.Executor
}

// NewService builds an ISaleFeatureService service.
func NewService(executor *request.Executor) *Service {
	return &Service{executor: executor}
}
