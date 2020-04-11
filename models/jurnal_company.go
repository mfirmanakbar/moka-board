package models

import (
	"github.com/jinzhu/gorm"
	"github.com/mfirmanakbar/moka-board/db"
)

// JurnalCompany has many JurnalUsers, JurnalCompanyID is the foreign key
type JurnalCompany struct {
	gorm.Model
	//Id              int64
	CompanyId   int64
	Email       string
	CompanyName string
	//CreatedAt       time.Time
	//UpdatedAt       time.Time
	//DeletedAt       time.Time
	JurnalUsers     []JurnalUser     `gorm:"foreignkey:JurnalCompanyId"`
	AccountMappings []AccountMapping `gorm:"foreignkey:JurnalCompanyId"`
}

func (j *JurnalCompany) FindConnectionsByCompanyId2(cid int64, showDeleted bool) (*[]Connection, error) {
	var err error
	_, _ = j.FindAccountMappingsByCompanyId(cid)
	var conns []Connection

	for index := range j.AccountMappings {
		if showDeleted {
			err = db.JurnalMokaGorm.Model(j.AccountMappings[index]).Unscoped().Related(&j.AccountMappings[index].Connections).Unscoped().Error
		} else {
			err = db.JurnalMokaGorm.Model(j.AccountMappings[index]).Related(&j.AccountMappings[index].Connections).Error
		}
	}
	if err != nil {
		return &conns, err
	}
	for _, ams := range j.AccountMappings {
		for _, cns := range ams.Connections {
			conns = append(conns, cns)
		}
	}
	return &conns, nil
}

func (j *JurnalCompany) FindConnectionsByCompanyId1(cid int64) (*JurnalCompany, error) {
	var err error
	_, _ = j.FindAccountMappingsByCompanyId(cid)
	for index, _ := range j.AccountMappings {
		err = db.JurnalMokaGorm.Model(j.AccountMappings[index]).Related(&j.AccountMappings[index].Connections).Unscoped().Error
	}
	if err != nil {
		return j, err
	}
	return j, nil
}

func (j *JurnalCompany) FindAccountMappingsByCompanyId(cid int64) (*JurnalCompany, error) {
	var err error
	_, _ = j.FindOneByCompanyId(cid)
	err = db.JurnalMokaGorm.Model(j).Related(&j.AccountMappings).Unscoped().Error
	if err != nil {
		return j, err
	}
	return j, nil
}

func (j *JurnalCompany) FindJurnalUsersByCompanyId(cid int64) (*JurnalCompany, error) {
	var err error
	_, _ = j.FindOneByCompanyId(cid)
	err = db.JurnalMokaGorm.Model(j).Related(&j.JurnalUsers).Error
	if err != nil {
		return j, err
	}
	return j, nil
}

// (*JurnalCompany, error) 	--> cuma type data yang harus di return
// (j *JurnalCompany) 		--> data object yang dipakai untuk simpen data dari DB
func (j *JurnalCompany) FindOneByCompanyId(cid int64) (*JurnalCompany, error) {
	var err error
	err = db.JurnalMokaGorm.Where("company_id = ?", cid).First(&j).Error
	if err != nil {
		return j, err
	}
	return j, nil
}

type JurnalCompanyInterface interface {
	FindOneByCompanyId(cid int64) (*JurnalCompany, error)
	FindJurnalUsersByCompanyId(cid int64) (*JurnalCompany, error)
	FindAccountMappingsByCompanyId(cid int64) (*JurnalCompany, error)
	FindConnectionsByCompanyId1(cid int64) (*JurnalCompany, error)
	FindConnectionsByCompanyId2(cid int64, showDeleted bool) (*[]Connection, error)
}

func JurnalCompanyDTO() JurnalCompanyInterface {
	return &JurnalCompany{}
}
