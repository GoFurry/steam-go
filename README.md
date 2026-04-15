# steam-go

`steam-go` is a lightweight Go SDK focused on the Steam Web API.

## Features

- Root `Client` with grouped service access under `client.API.*`
- Functional options for API key, access token, timeout, retry, rate limit, and proxy selection
- `key` and `access_token` are treated as different credentials and can be configured independently
- API key is optional and can be supplied through a rotating key provider
- Typed responses by default with matching raw response methods
- `401/429` can automatically retry with the next API key when `WithAPIKeys(...)` and `WithRetry(...)` are used together
- Shared request executor, centralized endpoint registry, and a single SDK error model
- No crawler, HTML parsing, A2S, heavy logging, or broad util package baggage

## Installation

```bash
go get github.com/GoFurry/steam-go@latest
```

## Go Version Policy

- `go.mod` is pinned to `Go 1.24`
- The project supports `Go 1.24+`
- CI continuously validates `Go 1.24` and `Go 1.25`

## Quick start

```go
package main

import (
	"context"
	"fmt"
	"time"

	steam "github.com/GoFurry/steam-go"
)

func main() {
	client, err := steam.NewClient(
		steam.WithTimeout(10*time.Second),
		steam.WithRetry(2),
	)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	resp, err := client.API.SteamUser.GetPlayerSummaries(
		context.Background(),
		[]string{"76561197960435530"},
	)
	if err != nil {
		panic(err)
	}

	for _, player := range resp.Response.Players {
		fmt.Printf("%s: %s\n", player.SteamID, player.PersonaName)
	}
}
```

## API overview

- `client.API.AccountCartService.GetCart(ctx, opts)`
- `client.API.AccountCartService.GetCartRaw(ctx, opts)`
- `client.API.AccountCartService.DeleteCart(ctx)`
- `client.API.AccountCartService.DeleteCartRaw(ctx)`
- `client.API.BillingService.GetRecurringSubscriptionsCount(ctx)`
- `client.API.BillingService.GetRecurringSubscriptionsCountRaw(ctx)`
- `client.API.CommunityService.GetApps(ctx, appIDs)`
- `client.API.CommunityService.GetAppsRaw(ctx, appIDs)`
- `client.API.FamilyGroupsService.GetChangeLog(ctx, familyGroupID)`
- `client.API.FamilyGroupsService.GetChangeLogRaw(ctx, familyGroupID)`
- `client.API.FamilyGroupsService.GetFamilyGroup(ctx, familyGroupID)`
- `client.API.FamilyGroupsService.GetFamilyGroupRaw(ctx, familyGroupID)`
- `client.API.FamilyGroupsService.GetFamilyGroupForUser(ctx, familyGroupID, opts)`
- `client.API.FamilyGroupsService.GetFamilyGroupForUserRaw(ctx, familyGroupID, opts)`
- `client.API.FamilyGroupsService.GetPlaytimeSummary(ctx, familyGroupID)`
- `client.API.FamilyGroupsService.GetPlaytimeSummaryRaw(ctx, familyGroupID)`
- `client.API.FamilyGroupsService.GetSharedLibraryApps(ctx, familyGroupID)`
- `client.API.FamilyGroupsService.GetSharedLibraryAppsRaw(ctx, familyGroupID)`
- `client.API.LoyaltyRewardsService.GetEquippedProfileItems(ctx, steamID, opts)`
- `client.API.LoyaltyRewardsService.GetEquippedProfileItemsRaw(ctx, steamID, opts)`
- `client.API.LoyaltyRewardsService.GetReactionsSummaryForUser(ctx, steamID)`
- `client.API.LoyaltyRewardsService.GetReactionsSummaryForUserRaw(ctx, steamID)`
- `client.API.LoyaltyRewardsService.GetSummary(ctx, steamID)`
- `client.API.LoyaltyRewardsService.GetSummaryRaw(ctx, steamID)`
- `client.API.SteamUser.GetPlayerSummaries(ctx, steamIDs)`
- `client.API.SteamUser.GetPlayerSummariesRaw(ctx, steamIDs)`
- `client.API.PlayerService.GetOwnedGames(ctx, steamID, opts)`
- `client.API.PlayerService.GetOwnedGamesRaw(ctx, steamID, opts)`
- `client.API.SteamNews.GetNewsForApp(ctx, appID, opts)`
- `client.API.SteamNews.GetNewsForAppRaw(ctx, appID, opts)`
- `client.API.SteamUserStats.GetPlayerAchievements(ctx, steamID, appID, opts)`
- `client.API.SteamUserStats.GetPlayerAchievementsRaw(ctx, steamID, appID, opts)`

## Options

- `WithAPIKey(key string)`
- `WithAPIKeys(keys ...string)`
- `WithAPIKeyProvider(provider APIKeyProvider)`
- `WithAccessToken(token string)`
- `WithAccessTokens(tokens ...string)`
- `WithAccessTokenProvider(provider AccessTokenProvider)`
- `WithBaseURL(url string)`
- `WithHTTPClient(client *http.Client)`
- `WithTimeout(timeout time.Duration)`
- `WithRetry(retry int)`
- `WithRateLimit(requestsPerSecond int)`
- `WithProxySelector(selector ProxySelector)`

## Error handling

SDK errors use `*steam.APIError` with these kinds:

- `request_build`
- `transport`
- `http_status`
- `decode`
- `api_response`

Use `errors.As(err, &apiErr)` to inspect kind, status code, and raw body.

## Examples

- [steamuser](examples/steamuser/main.go)
- [playerservice](examples/playerservice/main.go)
- [steamnews](examples/steamnews/main.go)
- [steamuserstats](examples/steamuserstats/main.go)
- Local manual scenario runner: `test/main.go`
