package steamchartsservice

import "github.com/gofurry/steam-go/internal/request"

// Service exposes ISteamChartsService methods.
type Service struct {
	executor *request.Executor
}

// NewService builds a SteamChartsService service.
func NewService(executor *request.Executor) *Service {
	return &Service{executor: executor}
}
