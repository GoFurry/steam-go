package userreviewsservice

import "github.com/gofurry/steam-go/internal/request"

// Service exposes IUserReviewsService methods.
type Service struct {
	executor *request.Executor
}

// NewService builds a UserReviewsService service.
func NewService(executor *request.Executor) *Service {
	return &Service{executor: executor}
}
