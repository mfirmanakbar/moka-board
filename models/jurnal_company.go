package models

import (
	"github.com/jinzhu/gorm"
	"github.com/mfirmanakbar/moka-board/datasources"
	"time"
)

// JurnalCompany has many JurnalUsers, JurnalCompanyID is the foreign key
type JurnalCompany struct {
	gorm.Model
	Id              int64
	CompanyId       int64
	Email           string
	CompanyName     string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
	JurnalUsers     []JurnalUser     `gorm:"foreignkey:JurnalCompanyId"`
	AccountMappings []AccountMapping `gorm:"foreignkey:JurnalCompanyId"`
}

func (j *JurnalCompany) FindConnectionsByCompanyId(cid int64) (*JurnalCompany, error) {
	var err error
	_, _ = j.FindAccountMappingsByCompanyId(cid)
	for index, _ := range j.AccountMappings {
		err = datasources.JmDb.Model(j.AccountMappings[index]).Related(&j.AccountMappings[index].Connections).Unscoped().Error
	}
	if err != nil {
		return j, err
	}
	return j, nil
}

func (j *JurnalCompany) FindAccountMappingsByCompanyId(cid int64) (*JurnalCompany, error) {
	var err error
	_, _ = j.FindOneByCompanyId(cid)
	err = datasources.JmDb.Model(j).Related(&j.AccountMappings).Unscoped().Error
	if err != nil {
		return j, err
	}
	return j, nil
}

func (j *JurnalCompany) FindJurnalUsersByCompanyId(cid int64) (*JurnalCompany, error) {
	var err error
	_, _ = j.FindOneByCompanyId(cid)
	err = datasources.JmDb.Model(j).Related(&j.JurnalUsers).Error
	if err != nil {
		return j, err
	}
	return j, nil
}

// (*JurnalCompany, error) 	--> cuma type data yang harus di return
// (j *JurnalCompany) 		--> data object yang dipakai untuk simpen data dari DB
func (j *JurnalCompany) FindOneByCompanyId(cid int64) (*JurnalCompany, error) {
	var err error
	err = datasources.JmDb.Where("company_id = ?", cid).First(&j).Error
	if err != nil {
		return j, err
	}
	return j, nil
}

type JurnalCompanyInterface interface {
	FindOneByCompanyId(cid int64) (*JurnalCompany, error)
	FindJurnalUsersByCompanyId(cid int64) (*JurnalCompany, error)
	FindAccountMappingsByCompanyId(cid int64) (*JurnalCompany, error)
	FindConnectionsByCompanyId(cid int64) (*JurnalCompany, error)
}

func JurnalCompanyDTO() JurnalCompanyInterface {
	return &JurnalCompany{}
}
