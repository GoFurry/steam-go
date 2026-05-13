package steamuserstats

import "github.com/gofurry/steam-go/internal/request"

// Service exposes ISteamUserStats methods.
type Service struct {
	executor *request.Executor
}

// NewService builds a SteamUserStats service.
func NewService(executor *request.Executor) *Service {
	return &Service{executor: executor}
}
