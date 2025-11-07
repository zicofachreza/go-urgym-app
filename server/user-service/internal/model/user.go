package model

import (
	"time"

	"github.com/zicofachreza/go-urgym-app/user-service/internal/utils"
	"gorm.io/gorm"
)

type User struct {
	ID                uint       `json:"id" gorm:"primaryKey"`
	Username          string     `json:"username" gorm:"unique" validate:"required,min=5"`
	Email             string     `json:"email" gorm:"unique" validate:"required,email"`
	Password          string     `json:"password" validate:"required,min=5"`
	IsMember          *bool      `json:"is_member" gorm:"default:false"`
	MembershipExpires *time.Time `json:"membership_expires"`
	ResetToken        *string    `json:"reset_token"`
	ResetTokenExpires *time.Time `json:"reset_token_expires"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

func (User) TableName() string {
	return "Users"
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if u.Password != "" {
		u.Password, err = utils.HashPassword(u.Password)
		if err != nil {
			return err
		}
	}
	return nil
}
