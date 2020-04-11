package models

import (
	"github.com/jinzhu/gorm"
	"github.com/mfirmanakbar/moka-board/db"
	"time"
)

type TransactionMapping struct {
	*gorm.Model
	//Id                  int64     `json:"id"`
	JurnalTransactionId int64     `json:"jurnal_transaction_id"`
	MokaTransactionId   int64     `json:"moka_transaction_id"`
	MokaTransactionType int       `json:"moka_transaction_type"`
	ConnectionId        int64     `json:"connection_id"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	DeletedAt           time.Time `json:"deleted_at"`
}

type TransactionParams struct {
	MokaTransactionType int
	ConnectionId        int64
	MokaTransactionId   int64
	JurnalTransactionId int64
}

type TransactionMappingInterface interface {
	SearchData(TransactionParams) (*[]TransactionMapping, error)
	QueryParams(TransactionParams) map[string]interface{}
}

func TransactionDTO() TransactionMappingInterface {
	return &TransactionMapping{}
}

func (tm *TransactionMapping) SearchData(prm TransactionParams) (*[]TransactionMapping, error) {
	var err error
	var transactionMappings []TransactionMapping

	if prm.ConnectionId < 1 {
		return &[]TransactionMapping{}, err
	}

	err = db.JurnalMokaGorm.Where(tm.QueryParams(prm)).Unscoped().Find(&transactionMappings).Error
	if err != nil {
		return &transactionMappings, err
	}
	return &transactionMappings, nil
}

func (tm *TransactionMapping) QueryParams(prm TransactionParams) map[string]interface{} {
	queryParams := make(map[string]interface{})
	if prm.MokaTransactionType > -1 {
		queryParams["moka_transaction_type"] = prm.MokaTransactionType
	}
	if prm.ConnectionId > 0 {
		queryParams["connection_id"] = prm.ConnectionId
	}
	if prm.MokaTransactionId > 0 {
		queryParams["moka_transaction_id"] = prm.MokaTransactionId
	}
	if prm.JurnalTransactionId > 0 {
		queryParams["jurnal_transaction_id"] = prm.JurnalTransactionId
	}
	return queryParams
}
