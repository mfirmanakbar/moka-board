package models

import (
	"github.com/mfirmanakbar/moka-board/datasources"
	"time"
)

type TransactionMapping struct {
	Id                  int64     `json:"id"`
	JurnalTransactionId int64     `json:"jurnal_transaction_id"`
	MokaTransactionId   int64     `json:"moka_transaction_id"`
	MokaTransactionType int       `json:"moka_transaction_type"`
	ConnectionId        int64     `json:"connection_id"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	DeletedAt           time.Time `json:"deleted_at"`
}

var (
	mysql = datasources.JmDb
)

type SearchParams struct {
	MokaTransactionType int   `gorm:"default:'-2'"`
	ConnectionId        int64 `gorm:"default:'-2'"`
	MokaTransactionId   int64 `gorm:"default:'-2'"`
	JurnalTransactionId int64 `gorm:"default:'-2'"`
}

func (tm *TransactionMapping) SearchTransactionMappings(prm SearchParams) (*[]TransactionMapping, error) {
	var err error
	var transactionMappings []TransactionMapping

	if prm.ConnectionId < 1 {
		return &[]TransactionMapping{}, err
	}

	err = datasources.JmDb.Where(queryParamsModified(prm)).Unscoped().Find(&transactionMappings).Error
	if err != nil {
		return &[]TransactionMapping{}, err
	}
	return &transactionMappings, nil
}

func queryParamsModified(prm SearchParams) map[string]interface{} {
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
