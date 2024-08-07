package rest

import "time"

type orderRequest struct {
	Products []string `json:"products"`
	Total    float32  `json:"total"`
	Status   string   `json:"status"`
}

type orderResponse struct {
	Id        string    `json:"_id"`
	Products  []string  `json:"products"`
	Total     float32   `json:"total"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
