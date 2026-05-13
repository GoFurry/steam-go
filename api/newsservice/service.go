package newsservice

import "github.com/gofurry/steam-go/internal/request"

// Service exposes INewsService methods.
type Service struct {
	executor *request.Executor
}

// NewService builds a NewsService service.
func NewService(executor *request.Executor) *Service {
	return &Service{executor: executor}
}
