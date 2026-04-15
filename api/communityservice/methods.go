package communityservice

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/GoFurry/steam-go/internal/endpoint"
	sdkerrors "github.com/GoFurry/steam-go/internal/errors"
	"github.com/GoFurry/steam-go/internal/request"
	"github.com/GoFurry/steam-go/internal/response"
)

// GetApps returns brief app metadata for the provided AppIDs.
func (s *Service) GetApps(ctx context.Context, appIDs []uint32) (GetAppsResponse, error) {
	body, err := s.GetAppsRaw(ctx, appIDs)
	if err != nil {
		return GetAppsResponse{}, err
	}
	return response.DecodeJSON[GetAppsResponse](body)
}

// GetAppsRaw returns the raw JSON response body.
func (s *Service) GetAppsRaw(ctx context.Context, appIDs []uint32) ([]byte, error) {
	if len(appIDs) == 0 {
		return nil, sdkerrors.New(sdkerrors.KindRequestBuild, 0, "at least one app id is required", nil, nil)
	}

	query := url.Values{}
	for idx, appID := range appIDs {
		if appID == 0 {
			return nil, sdkerrors.New(sdkerrors.KindRequestBuild, 0, "app id must be greater than zero", nil, nil)
		}
		query.Set("appids["+strconv.Itoa(idx)+"]", strconv.FormatUint(uint64(appID), 10))
	}

	return s.executor.DoRaw(ctx, request.RequestSpec{
		Method: http.MethodGet,
		Path:   endpoint.CommunityServiceGetApps,
		Query:  query,
	})
}
