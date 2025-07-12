package user

import "time"

type User struct {
	ID          string    `json:"id,omitempty" gorm:"primaryKey;type:text"` // UUID or custom string ID
	Name        string    `json:"name,omitempty" gorm:"not null"`
	Email       string    `json:"email,omitempty" gorm:"uniqueIndex;not null"`
	Password    string    `json:"-" gorm:"not null"` // hashed password, not exposed via JSON
	Phone       string    `json:"phone,omitempty" gorm:"not null"`
	AlternateNo string    `json:"alternateno,omitempty" gorm:""`
	Role        string    `json:"role,omitempty" gorm:"default:'buyer'"` // buyer, seller, admin
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	IsActive    bool      `json:"is_active,omitempty" gorm:"default:true"`
}
