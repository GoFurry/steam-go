package loyaltyrewardsservice

// GetEquippedProfileItemsResponse matches ILoyaltyRewardsService/GetEquippedProfileItems/v1.
type GetEquippedProfileItemsResponse struct {
	Response struct {
		ActiveDefinitions   []ProfileItemDefinition `json:"active_definitions"`
		InactiveDefinitions []ProfileItemDefinition `json:"inactive_definitions"`
	} `json:"response"`
}

// ProfileItemDefinition matches one Steam profile item definition.
type ProfileItemDefinition struct {
	AppID                 int64             `json:"appid"`
	DefID                 int64             `json:"defid"`
	Type                  int               `json:"type"`
	CommunityItemClass    int               `json:"community_item_class"`
	CommunityItemType     int               `json:"community_item_type"`
	PointCost             string            `json:"point_cost"`
	TimestampCreated      int64             `json:"timestamp_created"`
	TimestampUpdated      int64             `json:"timestamp_updated"`
	TimestampAvailable    int64             `json:"timestamp_available"`
	TimestampAvailableEnd int64             `json:"timestamp_available_end"`
	Quantity              string            `json:"quantity"`
	InternalDescription   string            `json:"internal_description"`
	Active                bool              `json:"active"`
	CommunityItemData     CommunityItemData `json:"community_item_data"`
	UsableDuration        int               `json:"usable_duration"`
	BundleDiscount        int               `json:"bundle_discount"`
}

// CommunityItemData contains profile item display data.
type CommunityItemData struct {
	ItemName        string `json:"item_name"`
	ItemTitle       string `json:"item_title"`
	ItemDescription string `json:"item_description"`
	ItemImageSmall  string `json:"item_image_small"`
	ItemImageLarge  string `json:"item_image_large"`
	ItemMovieWebm   string `json:"item_movie_webm"`
	ItemMovieMp4    string `json:"item_movie_mp4"`
	Animated        bool   `json:"animated"`
	Tiled           bool   `json:"tiled"`
}

// GetReactionsSummaryForUserResponse matches ILoyaltyRewardsService/GetReactionsSummaryForUser/v1.
type GetReactionsSummaryForUserResponse struct {
	Response struct {
		Total               []ReactionSummaryItem `json:"total"`
		UserReviews         []ReactionSummaryItem `json:"user_reviews"`
		UGC                 []ReactionSummaryItem `json:"ugc"`
		Profile             []ReactionSummaryItem `json:"profile"`
		TotalGiven          int64                 `json:"total_given"`
		TotalReceived       int64                 `json:"total_received"`
		TotalPointsGiven    string                `json:"total_points_given"`
		TotalPointsReceived string                `json:"total_points_received"`
	} `json:"response"`
}

// ReactionSummaryItem matches one reaction summary item.
type ReactionSummaryItem struct {
	ReactionID     int64  `json:"reactionid"`
	Given          int64  `json:"given"`
	Received       int64  `json:"received"`
	PointsGiven    string `json:"points_given"`
	PointsReceived string `json:"points_received"`
}

// GetSummaryResponse matches ILoyaltyRewardsService/GetSummary/v1.
type GetSummaryResponse struct {
	Response struct {
		Summary struct {
			Points       string `json:"points"`
			PointsEarned string `json:"points_earned"`
			PointsSpent  string `json:"points_spent"`
		} `json:"summary"`
		TimestampUpdated int64  `json:"timestamp_updated"`
		AuditIDHighwater string `json:"auditid_highwater"`
	} `json:"response"`
}
