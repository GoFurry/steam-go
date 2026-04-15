package communityservice

// GetAppsResponse matches ICommunityService/GetApps/v1.
type GetAppsResponse struct {
	Response struct {
		Apps []App `json:"apps"`
	} `json:"response"`
}

// App matches the community app metadata payload.
type App struct {
	AppID                            uint32 `json:"appid"`
	Name                             string `json:"name"`
	Icon                             string `json:"icon"`
	CommunityVisibleStats            bool   `json:"community_visible_stats,omitempty"`
	Propagation                      string `json:"propagation"`
	AppType                          int    `json:"app_type"`
	ContentDescriptorIDs             []int  `json:"content_descriptorids"`
	ContentDescriptorIDsIncludingDLC []int  `json:"content_descriptorids_including_dlc"`
}
