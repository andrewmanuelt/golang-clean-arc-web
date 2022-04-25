package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string
	Password   string
	Email      string
	Last_login time.Time
}
