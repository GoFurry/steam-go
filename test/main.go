package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	steam "github.com/GoFurry/steam-go"
	"github.com/GoFurry/steam-go/api/accountcartservice"
	"github.com/GoFurry/steam-go/api/familygroupsservice"
	"github.com/GoFurry/steam-go/api/loyaltyrewardsservice"
	"github.com/GoFurry/steam-go/api/playerservice"
	"github.com/GoFurry/steam-go/api/steamnews"
	"github.com/GoFurry/steam-go/api/steamuserstats"
)

const (
	defaultSteamID = "76561198370695025"
	defaultAppID   = uint32(550)
)

func main() {
	key := readCredential("key.txt")
	accessToken := readCredential("access-token.txt")
	familyGroupID := readCredential("family-group-id.txt")

	client, err := steam.NewClient(
		steam.WithAPIKey(key),
		steam.WithAccessToken(accessToken),
		steam.WithTimeout(30*time.Second),
		steam.WithRetry(1),
	)
	if err != nil {
		fatalf("create client failed: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	fmt.Println("== SteamUser.GetPlayerSummaries ==")
	playerResp, err := client.API.SteamUser.GetPlayerSummaries(ctx, []string{defaultSteamID})
	if err != nil {
		fatalf("GetPlayerSummaries failed: %v", err)
	}
	fmt.Printf("players=%d\n", len(playerResp.Response.Players))
	for _, player := range playerResp.Response.Players {
		fmt.Printf("steamid=%s persona=%s\n", player.SteamID, player.PersonaName)
	}

	fmt.Println("\n== PlayerService.GetOwnedGames ==")
	ownedGamesResp, err := client.API.PlayerService.GetOwnedGames(
		ctx,
		defaultSteamID,
		&playerservice.GetOwnedGamesOptions{IncludePlayedFreeGames: true},
	)
	if err != nil {
		fatalf("GetOwnedGames failed: %v", err)
	}
	fmt.Printf("game_count=%d\n", ownedGamesResp.Response.GameCount)
	for i, game := range ownedGamesResp.Response.Games {
		if i >= 10 {
			break
		}
		fmt.Printf("[%d] appid=%d name=%s playtime_forever=%d\n", i+1, game.AppID, game.Name, game.PlaytimeForever)
	}

	fmt.Println("\n== SteamNews.GetNewsForApp ==")
	newsResp, err := client.API.SteamNews.GetNewsForApp(
		ctx,
		defaultAppID,
		&steamnews.GetNewsForAppOptions{
			Count:     3,
			MaxLength: 200,
		},
	)
	if err != nil {
		fatalf("GetNewsForApp failed: %v", err)
	}
	fmt.Printf("news_count=%d\n", newsResp.AppNews.Count)
	for _, item := range newsResp.AppNews.NewsItems {
		fmt.Printf("gid=%s title=%s\n", item.GID, item.Title)
	}

	fmt.Println("\n== SteamUserStats.GetPlayerAchievements ==")
	achievementResp, err := client.API.SteamUserStats.GetPlayerAchievements(
		ctx,
		defaultSteamID,
		defaultAppID,
		&steamuserstats.GetPlayerAchievementsOptions{Language: "zh"},
	)
	if err != nil {
		fatalf("GetPlayerAchievements failed: %v", err)
	}
	fmt.Printf("achievement_count=%d\n", len(achievementResp.PlayerStats.Achievements))
	for i, achievement := range achievementResp.PlayerStats.Achievements {
		if i >= 10 {
			break
		}
		fmt.Printf("[%d] api=%s name=%s achieved=%d\n", i+1, achievement.APIName, achievement.Name, achievement.Achieved)
	}

	fmt.Println("\n== CommunityService.GetApps ==")
	appsResp, err := client.API.CommunityService.GetApps(ctx, []uint32{defaultAppID, 570})
	if err != nil {
		fatalf("GetApps failed: %v", err)
	}
	fmt.Printf("apps=%d\n", len(appsResp.Response.Apps))
	for _, app := range appsResp.Response.Apps {
		fmt.Printf("appid=%d name=%s\n", app.AppID, app.Name)
	}

	if accessToken == "" {
		fmt.Println("\n== Access token protected endpoints ==")
		fmt.Println("skip: access-token.txt is empty")
		return
	}

	fmt.Println("\n== AccountCartService.GetCart ==")
	cartResp, err := client.API.AccountCartService.GetCart(ctx, &accountcartservice.GetCartOptions{UserCountry: "cn"})
	if err != nil {
		fatalf("GetCart failed: %v", err)
	}
	fmt.Printf("cart_items=%d subtotal=%s\n", len(cartResp.Response.Cart.LineItems), cartResp.Response.Cart.Subtotal.FormattedAmount)

	fmt.Println("\n== BillingService.GetRecurringSubscriptionsCount ==")
	billingResp, err := client.API.BillingService.GetRecurringSubscriptionsCount(ctx)
	if err != nil {
		fatalf("GetRecurringSubscriptionsCount failed: %v", err)
	}
	fmt.Printf("active=%d inactive=%d\n", billingResp.Response.ActiveSubscriptionsCount, billingResp.Response.InactiveSubscriptionsCount)

	fmt.Println("\n== LoyaltyRewardsService.GetEquippedProfileItems ==")
	profileItemsResp, err := client.API.LoyaltyRewardsService.GetEquippedProfileItems(
		ctx,
		defaultSteamID,
		&loyaltyrewardsservice.GetEquippedProfileItemsOptions{Language: "zh"},
	)
	if err != nil {
		fatalf("GetEquippedProfileItems failed: %v", err)
	}
	fmt.Printf("active=%d inactive=%d\n", len(profileItemsResp.Response.ActiveDefinitions), len(profileItemsResp.Response.InactiveDefinitions))

	fmt.Println("\n== LoyaltyRewardsService.GetReactionsSummaryForUser ==")
	reactionsResp, err := client.API.LoyaltyRewardsService.GetReactionsSummaryForUser(ctx, defaultSteamID)
	if err != nil {
		fatalf("GetReactionsSummaryForUser failed: %v", err)
	}
	fmt.Printf("total_given=%d total_received=%d\n", reactionsResp.Response.TotalGiven, reactionsResp.Response.TotalReceived)

	fmt.Println("\n== LoyaltyRewardsService.GetSummary ==")
	summaryResp, err := client.API.LoyaltyRewardsService.GetSummary(ctx, defaultSteamID)
	if err != nil {
		fatalf("GetSummary failed: %v", err)
	}
	fmt.Printf("points=%s earned=%s spent=%s\n", summaryResp.Response.Summary.Points, summaryResp.Response.Summary.PointsEarned, summaryResp.Response.Summary.PointsSpent)

	if familyGroupID == "" {
		fmt.Println("\n== FamilyGroupsService.* ==")
		fmt.Println("skip: family-group-id.txt is empty")
		return
	}

	fmt.Println("\n== FamilyGroupsService.GetChangeLog ==")
	changeLogResp, err := client.API.FamilyGroupsService.GetChangeLog(ctx, familyGroupID)
	if err != nil {
		fatalf("GetChangeLog failed: %v", err)
	}
	fmt.Printf("changes=%d\n", len(changeLogResp.Response.Changes))

	fmt.Println("\n== FamilyGroupsService.GetFamilyGroup ==")
	familyGroupResp, err := client.API.FamilyGroupsService.GetFamilyGroup(ctx, familyGroupID)
	if err != nil {
		fatalf("GetFamilyGroup failed: %v", err)
	}
	fmt.Printf("name=%s members=%d\n", familyGroupResp.Response.Name, len(familyGroupResp.Response.Members))

	fmt.Println("\n== FamilyGroupsService.GetFamilyGroupForUser ==")
	familyForUserResp, err := client.API.FamilyGroupsService.GetFamilyGroupForUser(
		ctx,
		familyGroupID,
		&familygroupsservice.GetFamilyGroupForUserOptions{IncludeFamilyGroupResponse: true},
	)
	if err != nil {
		fatalf("GetFamilyGroupForUser failed: %v", err)
	}
	fmt.Printf("family_groupid=%s role=%d\n", familyForUserResp.Response.FamilyGroupID, familyForUserResp.Response.Role)

	fmt.Println("\n== FamilyGroupsService.GetPlaytimeSummary ==")
	playtimeResp, err := client.API.FamilyGroupsService.GetPlaytimeSummary(ctx, familyGroupID)
	if err != nil {
		fatalf("GetPlaytimeSummary failed: %v", err)
	}
	fmt.Printf("entries=%d\n", len(playtimeResp.Response.Entries))

	fmt.Println("\n== FamilyGroupsService.GetSharedLibraryApps ==")
	sharedAppsResp, err := client.API.FamilyGroupsService.GetSharedLibraryApps(ctx, familyGroupID)
	if err != nil {
		fatalf("GetSharedLibraryApps failed: %v", err)
	}
	fmt.Printf("owner=%s apps=%d\n", sharedAppsResp.Response.OwnerSteamID, len(sharedAppsResp.Response.Apps))
}

func readCredential(name string) string {
	body, err := os.ReadFile(name)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(body))
}

func fatalf(format string, args ...any) {
	fmt.Printf("ERROR: "+format+"\n", args...)
	os.Exit(1)
}
