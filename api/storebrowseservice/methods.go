package storebrowseservice

import (
	"context"
	"net/http"

	"github.com/gofurry/steam-go/internal/endpoint"
	"github.com/gofurry/steam-go/internal/request"
	"github.com/gofurry/steam-go/internal/response"
)

// GetContentHubConfig returns Steam content hub configuration metadata.
func (s *Service) GetContentHubConfig(ctx context.Context) (GetContentHubConfigResponse, error) {
	body, err := s.GetContentHubConfigRaw(ctx)
	if err != nil {
		return GetContentHubConfigResponse{}, err
	}
	return response.DecodeJSON[GetContentHubConfigResponse](body)
}

// GetContentHubConfigRaw returns the raw JSON response body.
func (s *Service) GetContentHubConfigRaw(ctx context.Context) ([]byte, error) {
	return s.executor.DoRaw(ctx, request.RequestSpec{
		Method: http.MethodGet,
		Path:   endpoint.StoreBrowseServiceGetContentHubConfig,
	})
}
