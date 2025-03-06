package entities

type Item struct {
	shortDescription string `json:"shortDescription"`
	price            string `json:"price"`
}

type Receipt struct {
	ID           string `json:"id"`
	retailer     string `json:"retailer"`
	purchaseDate string `json:"purchaseDate"`
	purchaseTime string `json:"purchaseTime"`
	items        []Item `json:"items"`
	total        string `json:"total"`
}
