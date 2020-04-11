package models

import (
	"github.com/mfirmanakbar/moka-board/db"
	"time"
)

type MokaAccount struct {
	//gorm.Model
	ID             int64
	UserId         int64
	Email          string
	Fullname       string
	CompanyName    string
	AccessToken    string
	RefreshToken   string
	CreatedAt      *time.Time `sql:"index" json:"created_at"`
	UpdatedAt      *time.Time `sql:"index" json:"updated_at"`
	DeletedAt      *time.Time `sql:"index" json:"deleted_at"`
	BusinessId     int64
	AccountMapping AccountMapping `gorm:"foreignkey:MokaUserId"`
}

func (m MokaAccount) FindByConnectionId(id int64) (*MokaAccount, error) {
	var err error
	var mokaAccount MokaAccount
	var connection Connection
	var accountMapping AccountMapping

	err = db.JurnalMokaGorm.First(&connection, id).Unscoped().Error
	if err != nil {
		return &mokaAccount, err
	}

	err = db.JurnalMokaGorm.First(&accountMapping, connection.AccountMappingId).Unscoped().Error
	if err != nil {
		return &mokaAccount, err
	}

	err = db.JurnalMokaGorm.First(&mokaAccount, accountMapping.MokaUserId).Unscoped().Error
	if err != nil {
		return &mokaAccount, err
	}

	return &mokaAccount, nil
}

type MokaAccountInterface interface {
	FindByConnectionId(id int64) (*MokaAccount, error)
}

func MokaAccountDTO() MokaAccountInterface {
	return &MokaAccount{}
}
