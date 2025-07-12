package review

import "time"

type Review struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	ProductID string    `json:"product_id"`
	Rating    int       `json:"rating"` // 1 to 5
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}
