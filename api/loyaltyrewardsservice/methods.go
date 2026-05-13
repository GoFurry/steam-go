package loyaltyrewardsservice

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/gofurry/steam-go/internal/endpoint"
	sdkerrors "github.com/gofurry/steam-go/internal/errors"
	"github.com/gofurry/steam-go/internal/request"
	"github.com/gofurry/steam-go/internal/response"
)

// GetEquippedProfileItemsOptions controls optional query parameters for GetEquippedProfileItems.
type GetEquippedProfileItemsOptions struct {
	Language string
}

// GetEquippedProfileItems returns profile item definitions for the given Steam ID.
func (s *Service) GetEquippedProfileItems(ctx context.Context, steamID string, opts *GetEquippedProfileItemsOptions) (GetEquippedProfileItemsResponse, error) {
	body, err := s.GetEquippedProfileItemsRaw(ctx, steamID, opts)
	if err != nil {
		return GetEquippedProfileItemsResponse{}, err
	}
	return response.DecodeJSON[GetEquippedProfileItemsResponse](body)
}

// GetEquippedProfileItemsRaw returns the raw JSON response body.
func (s *Service) GetEquippedProfileItemsRaw(ctx context.Context, steamID string, opts *GetEquippedProfileItemsOptions) ([]byte, error) {
	query, err := buildSteamIDQuery(steamID)
	if err != nil {
		return nil, err
	}
	if opts != nil {
		language := strings.TrimSpace(opts.Language)
		if language != "" {
			query.Set("language", language)
		}
	}

	return s.executor.DoRaw(ctx, request.RequestSpec{
		Method: http.MethodGet,
		Path:   endpoint.LoyaltyRewardsServiceGetEquippedProfileItems,
		Query:  query,
	})
}

// GetReactionsSummaryForUser returns reaction summary data for the given Steam ID.
func (s *Service) GetReactionsSummaryForUser(ctx context.Context, steamID string) (GetReactionsSummaryForUserResponse, error) {
	body, err := s.GetReactionsSummaryForUserRaw(ctx, steamID)
	if err != nil {
		return GetReactionsSummaryForUserResponse{}, err
	}
	return response.DecodeJSON[GetReactionsSummaryForUserResponse](body)
}

// GetReactionsSummaryForUserRaw returns the raw JSON response body.
func (s *Service) GetReactionsSummaryForUserRaw(ctx context.Context, steamID string) ([]byte, error) {
	query, err := buildSteamIDQuery(steamID)
	if err != nil {
		return nil, err
	}

	return s.executor.DoRaw(ctx, request.RequestSpec{
		Method: http.MethodGet,
		Path:   endpoint.LoyaltyRewardsServiceGetReactionsSummaryForUser,
		Query:  query,
	})
}

// GetSummary returns loyalty points summary for the given Steam ID.
func (s *Service) GetSummary(ctx context.Context, steamID string) (GetSummaryResponse, error) {
	body, err := s.GetSummaryRaw(ctx, steamID)
	if err != nil {
		return GetSummaryResponse{}, err
	}
	return response.DecodeJSON[GetSummaryResponse](body)
}

// GetSummaryRaw returns the raw JSON response body.
func (s *Service) GetSummaryRaw(ctx context.Context, steamID string) ([]byte, error) {
	query, err := buildSteamIDQuery(steamID)
	if err != nil {
		return nil, err
	}

	return s.executor.DoRaw(ctx, request.RequestSpec{
		Method: http.MethodGet,
		Path:   endpoint.LoyaltyRewardsServiceGetSummary,
		Query:  query,
	})
}

func buildSteamIDQuery(steamID string) (url.Values, error) {
	trimmed := strings.TrimSpace(steamID)
	if trimmed == "" {
		return nil, sdkerrors.New(sdkerrors.KindRequestBuild, 0, "steam id must not be empty", nil, nil)
	}

	query := url.Values{}
	query.Set("steamid", trimmed)
	return query, nil
}
