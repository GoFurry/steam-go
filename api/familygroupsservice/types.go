package familygroupsservice

// GetChangeLogResponse matches IFamilyGroupsService/GetChangeLog/v1.
type GetChangeLogResponse struct {
	Response struct {
		Changes []FamilyGroupChange `json:"changes"`
	} `json:"response"`
}

// FamilyGroupChange matches one family group change log entry.
type FamilyGroupChange struct {
	Timestamp    string `json:"timestamp"`
	ActorSteamID string `json:"actor_steamid"`
	Type         int    `json:"type"`
	Body         string `json:"body"`
	BySupport    bool   `json:"by_support"`
}

// GetFamilyGroupResponse matches IFamilyGroupsService/GetFamilyGroup/v1.
type GetFamilyGroupResponse struct {
	Response FamilyGroup `json:"response"`
}

// FamilyGroup matches the family group payload.
type FamilyGroup struct {
	Name                         string              `json:"name"`
	Members                      []FamilyGroupMember `json:"members"`
	FreeSpots                    int                 `json:"free_spots"`
	Country                      string              `json:"country"`
	SlotCooldownRemainingSeconds int64               `json:"slot_cooldown_remaining_seconds"`
	SlotCooldownOverrides        int                 `json:"slot_cooldown_overrides"`
}

// FamilyGroupMember matches one family group member.
type FamilyGroupMember struct {
	SteamID                  string `json:"steamid"`
	Role                     int    `json:"role"`
	TimeJoined               int64  `json:"time_joined"`
	CooldownSecondsRemaining int64  `json:"cooldown_seconds_remaining"`
}

// GetFamilyGroupForUserResponse matches IFamilyGroupsService/GetFamilyGroupForUser/v1.
type GetFamilyGroupForUserResponse struct {
	Response struct {
		FamilyGroupID               string                         `json:"family_groupid"`
		IsNotMemberOfAnyGroup       bool                           `json:"is_not_member_of_any_group"`
		LatestTimeJoined            int64                          `json:"latest_time_joined"`
		LatestJoinedFamilyGroupID   string                         `json:"latest_joined_family_groupid"`
		Role                        int                            `json:"role"`
		CooldownSecondsRemaining    int64                          `json:"cooldown_seconds_remaining"`
		FamilyGroup                 FamilyGroup                    `json:"family_group"`
		CanUndeleteLastJoinedFamily bool                           `json:"can_undelete_last_joined_family"`
		MembershipHistory           []FamilyGroupMembershipHistory `json:"membership_history"`
	} `json:"response"`
}

// FamilyGroupMembershipHistory matches one family group membership history item.
type FamilyGroupMembershipHistory struct {
	FamilyGroupID string `json:"family_groupid"`
	RTimeJoined   int64  `json:"rtime_joined"`
	RTimeLeft     int64  `json:"rtime_left"`
	Role          int    `json:"role"`
	Participated  bool   `json:"participated"`
}

// GetPlaytimeSummaryResponse matches IFamilyGroupsService/GetPlaytimeSummary/v1.
type GetPlaytimeSummaryResponse struct {
	Response struct {
		Entries []PlaytimeEntry `json:"entries"`
	} `json:"response"`
}

// PlaytimeEntry matches one family playtime record.
type PlaytimeEntry struct {
	SteamID       string `json:"steamid"`
	AppID         uint64 `json:"appid"`
	FirstPlayed   int64  `json:"first_played"`
	LatestPlayed  int64  `json:"latest_played"`
	SecondsPlayed int64  `json:"seconds_played"`
}

// GetSharedLibraryAppsResponse matches IFamilyGroupsService/GetSharedLibraryApps/v1.
type GetSharedLibraryAppsResponse struct {
	Response struct {
		OwnerSteamID string             `json:"owner_steamid"`
		Apps         []SharedLibraryApp `json:"apps"`
	} `json:"response"`
}

// SharedLibraryApp matches one family shared app payload.
type SharedLibraryApp struct {
	AppID           uint64   `json:"appid"`
	OwnerSteamIDs   []string `json:"owner_steamids"`
	Name            string   `json:"name"`
	CapsuleFilename string   `json:"capsule_filename"`
	ImgIconHash     string   `json:"img_icon_hash"`
	ExcludeReason   int      `json:"exclude_reason"`
	RtTimeAcquired  int64    `json:"rt_time_acquired"`
	RtLastPlayed    int64    `json:"rt_last_played"`
	RtPlaytime      int64    `json:"rt_playtime"`
	AppType         int      `json:"app_type"`
}
