package billingservice

import "github.com/gofurry/steam-go/internal/request"

// Service exposes IBillingService methods.
type Service struct {
	executor *request.Executor
}

// NewService builds an IBillingService service.
func NewService(executor *request.Executor) *Service {
	return &Service{executor: executor}
}
