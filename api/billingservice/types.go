package billingservice

// GetRecurringSubscriptionsCountResponse matches IBillingService/GetRecurringSubscriptionsCount/v1.
type GetRecurringSubscriptionsCountResponse struct {
	Response struct {
		ActiveSubscriptionsCount   int64 `json:"active_subscriptions_count"`
		InactiveSubscriptionsCount int64 `json:"inactive_subscriptions_count"`
	} `json:"response"`
}
