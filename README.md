# steam-go

`steam-go` is a lightweight Go SDK focused on the Steam Web API.

## Features

- Root `Client` with service-oriented access to `AccountCartService`, `BillingService`, `CommunityService`, `FamilyGroupsService`, `LoyaltyRewardsService`, `SteamUser`, `PlayerService`, `SteamNews`, and `SteamUserStats`
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

	resp, err := client.SteamUser.GetPlayerSummaries(
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

- `client.AccountCartService.GetCart(ctx, opts)`
- `client.AccountCartService.GetCartRaw(ctx, opts)`
- `client.AccountCartService.DeleteCart(ctx)`
- `client.AccountCartService.DeleteCartRaw(ctx)`
- `client.BillingService.GetRecurringSubscriptionsCount(ctx)`
- `client.BillingService.GetRecurringSubscriptionsCountRaw(ctx)`
- `client.CommunityService.GetApps(ctx, appIDs)`
- `client.CommunityService.GetAppsRaw(ctx, appIDs)`
- `client.FamilyGroupsService.GetChangeLog(ctx, familyGroupID)`
- `client.FamilyGroupsService.GetChangeLogRaw(ctx, familyGroupID)`
- `client.FamilyGroupsService.GetFamilyGroup(ctx, familyGroupID)`
- `client.FamilyGroupsService.GetFamilyGroupRaw(ctx, familyGroupID)`
- `client.FamilyGroupsService.GetFamilyGroupForUser(ctx, familyGroupID, opts)`
- `client.FamilyGroupsService.GetFamilyGroupForUserRaw(ctx, familyGroupID, opts)`
- `client.FamilyGroupsService.GetPlaytimeSummary(ctx, familyGroupID)`
- `client.FamilyGroupsService.GetPlaytimeSummaryRaw(ctx, familyGroupID)`
- `client.FamilyGroupsService.GetSharedLibraryApps(ctx, familyGroupID)`
- `client.FamilyGroupsService.GetSharedLibraryAppsRaw(ctx, familyGroupID)`
- `client.LoyaltyRewardsService.GetEquippedProfileItems(ctx, steamID, opts)`
- `client.LoyaltyRewardsService.GetEquippedProfileItemsRaw(ctx, steamID, opts)`
- `client.LoyaltyRewardsService.GetReactionsSummaryForUser(ctx, steamID)`
- `client.LoyaltyRewardsService.GetReactionsSummaryForUserRaw(ctx, steamID)`
- `client.LoyaltyRewardsService.GetSummary(ctx, steamID)`
- `client.LoyaltyRewardsService.GetSummaryRaw(ctx, steamID)`
- `client.SteamUser.GetPlayerSummaries(ctx, steamIDs)`
- `client.SteamUser.GetPlayerSummariesRaw(ctx, steamIDs)`
- `client.PlayerService.GetOwnedGames(ctx, steamID, opts)`
- `client.PlayerService.GetOwnedGamesRaw(ctx, steamID, opts)`
- `client.SteamNews.GetNewsForApp(ctx, appID, opts)`
- `client.SteamNews.GetNewsForAppRaw(ctx, appID, opts)`
- `client.SteamUserStats.GetPlayerAchievements(ctx, steamID, appID, opts)`
- `client.SteamUserStats.GetPlayerAchievementsRaw(ctx, steamID, appID, opts)`

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
