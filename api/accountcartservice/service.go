package accountcartservice

import "github.com/gofurry/steam-go/internal/request"

// Service exposes IAccountCartService methods.
type Service struct {
	executor *request.Executor
}

// NewService builds an IAccountCartService service.
func NewService(executor *request.Executor) *Service {
	return &Service{executor: executor}
}
