package models

type Event struct {
	Context Context `json:"context"`

	// TODO: Finalise the track event feilds
	Properties struct {
		// Feilds of Category, Add/Remove Product
		Category  string  `json:"category"`
		ProductID string  `json:"id"`
		SKU       string  `json:"sku"`
		Name      string  `json:"name"`
		Price     float32 `json:"price"`
		Quantity  int     `json:"quantity"`

		// Feilds for Complete Cart
		OrderID  string    `json:"orderId" cql:"orderid"`
		Total    float32   `json:"total" cql:"total"`
		Products []Product `json:"products" cql:"products"`

		// CQL properties
		Cost float32 `json:"cost"`
	} `json:"properties"`

	Event       string `json:"event"`
	UserID      string `json:"userId"`
	AnonymousID string `json:"anonymousId"`
	// TODO: Check for a timing lib for go
	Type      string `json:"type"`
	Timestamp string `json:"timestamp"`
	AppID     string `json:"appId"`
	MessageID string `json:"messageId"`
	SentAt    string `json:"sentAt"`
}
