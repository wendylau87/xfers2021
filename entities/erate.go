package entities

type ERate struct {
	ID        int     `json:"id"`
	KursID    int     `json:"kurs_id"`
	Buy       float64 `json:"buy"`
	Sell      float64 `json:"sell"`
	ValidDate string  `json:"valid_date"`
}

type CreateERate struct {
	Buy  float64 `json:"beli"`
	Sell float64 `json:"jual"`
}

type ERateResponse struct {
	Buy       float64 `json:"beli"`
	Sell      float64 `json:"jual"`
	ValidDate string  `json:"date"`
}
