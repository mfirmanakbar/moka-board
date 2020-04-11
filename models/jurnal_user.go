package models

import (
	"github.com/mfirmanakbar/moka-board/db"
	"time"
)

type JurnalUser struct {
	//gorm.Model
	ID              int64 `json:"id"`
	JurnalCompanyId int64
	Email           string
	AccessToken     string
	CreatedAt       *time.Time `sql:"index" json:"created_at"`
	UpdatedAt       *time.Time `sql:"index" json:"updated_at"`
	DeletedAt       *time.Time `sql:"index" json:"deleted_at"`
	ViewSyncGuide   int
	ViewStepGuide   int
}

func (j JurnalUser) FindByConnectionId(id int64) (*[]JurnalUser, error) {
	var err error
	var jurnalUsers []JurnalUser
	var connection Connection
	var accountMapping AccountMapping

	err = db.JurnalMokaGorm.First(&connection, id).Unscoped().Error
	if err != nil {
		return &jurnalUsers, err
	}
	err = db.JurnalMokaGorm.First(&accountMapping, connection.AccountMappingId).Unscoped().Error
	if err != nil {
		return &jurnalUsers, err
	}
	err = db.JurnalMokaGorm.Where("jurnal_company_id = ?", accountMapping.JurnalCompanyId).Find(&jurnalUsers).Unscoped().Error
	if err != nil {
		return &jurnalUsers, err
	}

	return &jurnalUsers, nil
}

func (j JurnalUser) FindOne(id int64) (*JurnalUser, error) {
	var err error
	var jurnalUser JurnalUser
	err = db.JurnalMokaGorm.First(&jurnalUser, id).Error
	if err != nil {
		return &jurnalUser, err
	}
	return &jurnalUser, nil
}

type JurnalUserInterface interface {
	FindOne(id int64) (*JurnalUser, error)
	FindByConnectionId(id int64) (*[]JurnalUser, error)
}

func JurnalUserDTO() JurnalUserInterface {
	return &JurnalUser{}
}
