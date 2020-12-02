package pkg

type OrderCreatedEvent struct {
	OrderCode string `json:"orderCode"`
}

type CartEntry struct {
	Quantity int64             `json:"quantity"`
	Product  map[string]string `json:"product"`
}
