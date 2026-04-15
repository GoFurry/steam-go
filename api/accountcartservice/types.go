package accountcartservice

// GetCartResponse matches IAccountCartService/GetCart/v1.
type GetCartResponse struct {
	Response struct {
		Cart Cart `json:"cart"`
	} `json:"response"`
}

// DeleteCartResponse matches IAccountCartService/DeleteCart/v1.
type DeleteCartResponse struct {
	Response struct{} `json:"response"`
}

// Cart is the cart payload.
type Cart struct {
	LineItems []LineItem  `json:"line_items"`
	Subtotal  MoneyAmount `json:"subtotal"`
	IsValid   bool        `json:"is_valid"`
}

// LineItem matches a Steam cart line item.
type LineItem struct {
	LineItemID     string        `json:"line_item_id"`
	Type           int           `json:"type"`
	PackageID      uint32        `json:"packageid"`
	IsValid        bool          `json:"is_valid"`
	TimeAdded      int64         `json:"time_added"`
	PriceWhenAdded MoneyAmount   `json:"price_when_added"`
	Flags          LineItemFlags `json:"flags"`
}

// LineItemFlags contains cart line item flags.
type LineItemFlags struct {
	IsGift    bool `json:"is_gift"`
	IsPrivate bool `json:"is_private"`
}

// MoneyAmount matches Steam money fields encoded in cart payloads.
type MoneyAmount struct {
	AmountInCents   string `json:"amount_in_cents"`
	CurrencyCode    int    `json:"currency_code"`
	FormattedAmount string `json:"formatted_amount"`
}
