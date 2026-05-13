package storebrowseservice

import "github.com/gofurry/steam-go/internal/request"

// Service exposes IStoreBrowseService methods.
type Service struct {
	executor *request.Executor
}

// NewService builds a StoreBrowseService service.
func NewService(executor *request.Executor) *Service {
	return &Service{executor: executor}
}
