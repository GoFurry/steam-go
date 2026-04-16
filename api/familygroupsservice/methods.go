package familygroupsservice

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/GoFurry/steam-go/internal/endpoint"
	sdkerrors "github.com/GoFurry/steam-go/internal/errors"
	"github.com/GoFurry/steam-go/internal/request"
	"github.com/GoFurry/steam-go/internal/response"
)

// GetFamilyGroupForUserOptions controls optional query parameters for GetFamilyGroupForUser.
type GetFamilyGroupForUserOptions struct {
	IncludeFamilyGroupResponse bool
}

// GetChangeLog returns family group change logs.
func (s *Service) GetChangeLog(ctx context.Context, familyGroupID string) (GetChangeLogResponse, error) {
	body, err := s.GetChangeLogRaw(ctx, familyGroupID)
	if err != nil {
		return GetChangeLogResponse{}, err
	}
	return response.DecodeJSON[GetChangeLogResponse](body)
}

// GetChangeLogRaw returns the raw JSON response body.
func (s *Service) GetChangeLogRaw(ctx context.Context, familyGroupID string) ([]byte, error) {
	return s.doWithFamilyGroupID(ctx, http.MethodGet, endpoint.FamilyGroupsServiceGetChangeLog, familyGroupID, nil)
}

// GetFamilyGroup returns family group details.
func (s *Service) GetFamilyGroup(ctx context.Context, familyGroupID string) (GetFamilyGroupResponse, error) {
	body, err := s.GetFamilyGroupRaw(ctx, familyGroupID)
	if err != nil {
		return GetFamilyGroupResponse{}, err
	}
	return response.DecodeJSON[GetFamilyGroupResponse](body)
}

// GetFamilyGroupRaw returns the raw JSON response body.
func (s *Service) GetFamilyGroupRaw(ctx context.Context, familyGroupID string) ([]byte, error) {
	return s.doWithFamilyGroupID(ctx, http.MethodGet, endpoint.FamilyGroupsServiceGetFamilyGroup, familyGroupID, nil)
}

// GetFamilyGroupForUser returns the current user's family group state.
func (s *Service) GetFamilyGroupForUser(ctx context.Context, familyGroupID string, opts *GetFamilyGroupForUserOptions) (GetFamilyGroupForUserResponse, error) {
	body, err := s.GetFamilyGroupForUserRaw(ctx, familyGroupID, opts)
	if err != nil {
		return GetFamilyGroupForUserResponse{}, err
	}
	return response.DecodeJSON[GetFamilyGroupForUserResponse](body)
}

// GetFamilyGroupForUserRaw returns the raw JSON response body.
func (s *Service) GetFamilyGroupForUserRaw(ctx context.Context, familyGroupID string, opts *GetFamilyGroupForUserOptions) ([]byte, error) {
	query := url.Values{}
	if opts != nil && opts.IncludeFamilyGroupResponse {
		query.Set("include_family_group_response", "true")
	}
	return s.doWithFamilyGroupID(ctx, http.MethodGet, endpoint.FamilyGroupsServiceGetFamilyGroupForUser, familyGroupID, query)
}

// GetPlaytimeSummary returns family group playtime data.
func (s *Service) GetPlaytimeSummary(ctx context.Context, familyGroupID string) (GetPlaytimeSummaryResponse, error) {
	body, err := s.GetPlaytimeSummaryRaw(ctx, familyGroupID)
	if err != nil {
		return GetPlaytimeSummaryResponse{}, err
	}
	return response.DecodeJSON[GetPlaytimeSummaryResponse](body)
}

// GetPlaytimeSummaryRaw returns the raw JSON response body.
func (s *Service) GetPlaytimeSummaryRaw(ctx context.Context, familyGroupID string) ([]byte, error) {
	return s.doWithFamilyGroupID(ctx, http.MethodPost, endpoint.FamilyGroupsServiceGetPlaytimeSummary, familyGroupID, nil)
}

// GetSharedLibraryApps returns family shared library apps.
func (s *Service) GetSharedLibraryApps(ctx context.Context, familyGroupID string) (GetSharedLibraryAppsResponse, error) {
	body, err := s.GetSharedLibraryAppsRaw(ctx, familyGroupID)
	if err != nil {
		return GetSharedLibraryAppsResponse{}, err
	}
	return response.DecodeJSON[GetSharedLibraryAppsResponse](body)
}

// GetSharedLibraryAppsRaw returns the raw JSON response body.
func (s *Service) GetSharedLibraryAppsRaw(ctx context.Context, familyGroupID string) ([]byte, error) {
	return s.doWithFamilyGroupID(ctx, http.MethodGet, endpoint.FamilyGroupsServiceGetSharedLibraryApps, familyGroupID, nil)
}

func (s *Service) doWithFamilyGroupID(ctx context.Context, method string, path string, familyGroupID string, query url.Values) ([]byte, error) {
	trimmed := strings.TrimSpace(familyGroupID)
	if trimmed == "" {
		return nil, sdkerrors.New(sdkerrors.KindRequestBuild, 0, "family group id must not be empty", nil, nil)
	}
	if query == nil {
		query = url.Values{}
	}
	query.Set("family_groupid", trimmed)

	return s.executor.DoRaw(ctx, request.RequestSpec{
		Method: method,
		Path:   path,
		Query:  query,
	})
}
