package familygroupsservice

import "github.com/gofurry/steam-go/internal/request"

// Service exposes IFamilyGroupsService methods.
type Service struct {
	executor *request.Executor
}

// NewService builds an IFamilyGroupsService service.
func NewService(executor *request.Executor) *Service {
	return &Service{executor: executor}
}
