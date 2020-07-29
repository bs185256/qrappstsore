package items

type catalogItem struct {
	ShortDescription struct {
		Values []struct {
			Locale string `json:"locale,omitempty"`
			Value  string `json:"value,omitempty"`
		} `json:"values,omitempty"`
	} `json:"shortDescription,omitempty"`
	ItemPrices []struct {
		Price    float64 `json:"price,omitempty"`
		Currency string  `json:"currency,omitempty"`
		Status   string  `json:"status,omitempty"`
	} `json:"itemPrices,omitempty"`
}

type items struct {
	Items []*item `json:"items,omitempty"`
}

type item struct {
	Name  string  `json:"name,omitempty"`
	Price float64 `json:"price,omitempty"`
}
