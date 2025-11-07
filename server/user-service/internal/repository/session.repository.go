package repository

import (
	"gorm.io/gorm"
)

type SessionRepository struct {
	DB *gorm.DB
}
