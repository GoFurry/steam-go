package steamnotificationservice

import "github.com/gofurry/steam-go/internal/request"

// Service exposes ISteamNotificationService methods.
type Service struct {
	executor *request.Executor
}

// NewService builds a SteamNotificationService service.
func NewService(executor *request.Executor) *Service {
	return &Service{executor: executor}
}
