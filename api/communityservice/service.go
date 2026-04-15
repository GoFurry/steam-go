package communityservice

import "github.com/GoFurry/steam-go/internal/request"

// Service exposes ICommunityService methods.
type Service struct {
	executor *request.Executor
}

// NewService builds an ICommunityService service.
func NewService(executor *request.Executor) *Service {
	return &Service{executor: executor}
}
