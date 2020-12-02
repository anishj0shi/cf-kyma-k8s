package pkg

type OrderCreatedEvent struct {
	OrderCode string `json:"orderCode"`
}

type CartEntry struct {
	Quantity int64   `json:"quantity"`
	Product  Product `json:"product"`
}

type Product struct {
	Code string `json:"code"`
}
