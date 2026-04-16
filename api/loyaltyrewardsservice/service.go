package loyaltyrewardsservice

import "github.com/GoFurry/steam-go/internal/request"

// Service exposes ILoyaltyRewardsService methods.
type Service struct {
	executor *request.Executor
}

// NewService builds an ILoyaltyRewardsService service.
func NewService(executor *request.Executor) *Service {
	return &Service{executor: executor}
}
