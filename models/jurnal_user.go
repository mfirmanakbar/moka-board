package models

import (
	"github.com/jinzhu/gorm"
	"github.com/mfirmanakbar/moka-board/datasources"
	"time"
)

type JurnalUser struct {
	gorm.Model
	Id              int64
	JurnalCompanyId int64
	Email           string
	AccessToken     string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
	ViewSyncGuide   int
	ViewStepGuide   int
}

func (j JurnalUser) FindOne(id int64) (*JurnalUser, error) {
	var err error
	var jurnalUser JurnalUser
	err = datasources.JmDb.First(&jurnalUser, id).Error
	if err != nil {
		return &jurnalUser, err
	}
	return &jurnalUser, nil
}

type JurnalUserInterface interface {
	FindOne(id int64) (*JurnalUser, error)
}

func JurnalUserDTO() JurnalUserInterface {
	return &JurnalUser{}
}
