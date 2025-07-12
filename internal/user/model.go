package user

import "time"

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`    // hashed password
	Role      string    `json:"role"` // "buyer", "seller", "admin"
	CreatedAt time.Time `json:"created_at"`
}
