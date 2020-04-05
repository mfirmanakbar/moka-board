package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type MokaAccount struct {
	gorm.Model
	Id             int64
	UserId         int64
	Email          string
	Fullname       string
	CompanyName    string
	AccessToken    string
	RefreshToken   string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
	BusinessId     int64
	AccountMapping AccountMapping `gorm:"foreignkey:MokaUserId"`
}
