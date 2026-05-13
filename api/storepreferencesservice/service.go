package storepreferencesservice

import "github.com/gofurry/steam-go/internal/request"

// Service exposes IStorePreferencesService methods.
type Service struct {
	executor *request.Executor
}

// NewService builds a StorePreferencesService service.
func NewService(executor *request.Executor) *Service {
	return &Service{executor: executor}
}
