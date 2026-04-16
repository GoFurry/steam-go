package billingservice

import (
	"context"
	"net/http"

	"github.com/GoFurry/steam-go/internal/endpoint"
	"github.com/GoFurry/steam-go/internal/request"
	"github.com/GoFurry/steam-go/internal/response"
)

// GetRecurringSubscriptionsCount returns subscription counts for the caller.
func (s *Service) GetRecurringSubscriptionsCount(ctx context.Context) (GetRecurringSubscriptionsCountResponse, error) {
	body, err := s.GetRecurringSubscriptionsCountRaw(ctx)
	if err != nil {
		return GetRecurringSubscriptionsCountResponse{}, err
	}
	return response.DecodeJSON[GetRecurringSubscriptionsCountResponse](body)
}

// GetRecurringSubscriptionsCountRaw returns the raw JSON response body.
func (s *Service) GetRecurringSubscriptionsCountRaw(ctx context.Context) ([]byte, error) {
	return s.executor.DoRaw(ctx, request.RequestSpec{
		Method: http.MethodGet,
		Path:   endpoint.BillingServiceGetRecurringSubscriptionsCount,
	})
}
