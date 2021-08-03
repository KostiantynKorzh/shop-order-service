package model

type FullOrderInfo struct {
	Title     string  `json:"title"`
	Price     float64 `json:"price"`
	ImagePath string  `json:"imagePath"`
	Quantity  uint    `json:"quantity"`
}

func (o *FullOrderInfo) SetQuantity(quantity uint) {
	o.Quantity = quantity
}
