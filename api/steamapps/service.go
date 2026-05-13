package steamapps

import "github.com/gofurry/steam-go/internal/request"

// Service exposes ISteamApps methods.
type Service struct {
	executor *request.Executor
}

// NewService builds a SteamApps service.
func NewService(executor *request.Executor) *Service {
	return &Service{executor: executor}
}
