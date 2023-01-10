package types

type AddProductBody struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Category int    `json:"category"`
	Quantity int    `json:"quantity"`
	Picture  string `json:"picture"`
}

type UpdateProductBody struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Category int    `json:"category"`
	Quantity int    `json:"quantity"`
}
