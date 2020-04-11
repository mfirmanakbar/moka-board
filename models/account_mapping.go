package models

import (
	"time"
)

// AccountMapping has many Connections, AccountMappingID (on Connection) is the foreign key
type AccountMapping struct {
	Id              int64 `gorm:"primary_key"`
	JurnalCompanyId int64
	MokaUserId      int64
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Connections     []Connection `gorm:"foreignkey:AccountMappingId"`
}
