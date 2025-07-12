package payment

import "time"

type Payment struct {
	ID        string    `json:"id"`
	OrderID   string    `json:"order_id"`
	UserID    string    `json:"user_id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`   // pending, success, failed
	Provider  string    `json:"provider"` // Razorpay, Stripe, etc.
	CreatedAt time.Time `json:"created_at"`
}
