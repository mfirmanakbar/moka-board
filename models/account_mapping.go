package models

import (
	"time"
)

type AccountMapping struct {
	Id              int64 `gorm:"primary_key"`
	JurnalCompanyId int64
	MokaUserId      int64
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Connections     []Connection `gorm:"foreignkey:AccountMappingId"`
}
