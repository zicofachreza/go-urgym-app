package model

import "time"

type Session struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserID      uint      `json:"user_id"`
	HashedToken string    `json:"hashed_token"`
	DeviceInfo  string    `json:"device_info"`
	IpAddress   string    `json:"ip_address"`
	ExpiresAt   time.Time `json:"expires_at"`
	LastUsedAt  time.Time `json:"last_used_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Session) TableName() string {
	return "Sessions"
}
