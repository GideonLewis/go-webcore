package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64
	FirstName string
	LastName  string
	Age       int64
	Gender    string
	Phone     string
	Email     string
	Country   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  *gorm.DeletedAt
}
