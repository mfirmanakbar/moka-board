package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TransactionMapping struct {
	Id                  int64     `json:"id"`
	JurnalTransactionId int64     `json:"jurnal_transaction_id"`
	MokaTransactionId   int64     `json:"moka_transaction_id"`
	MokaTransactionType int8      `json:"moka_transaction_type"`
	ConnectionId        int64     `json:"connection_id"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	DeletedAt           time.Time `json:"deleted_at"`
}

type TransactionMappings []*TransactionMapping

// (tm *TransactionMappingList) 		=> parameter from UI that called this method
// (*TransactionMappingList, error) 	=> return list from data and error
func (tm *TransactionMapping) SearchTransactionMappings(db *gorm.DB, conId int64) (*[]TransactionMapping, error) {
	var err error
	var transactionMappings []TransactionMapping
	//transactionMappings := []TransactionMapping{}
	err = db.Debug().Model(&TransactionMapping{}).Limit(10).Find(&transactionMappings).Error
	if err != nil {
		return &[]TransactionMapping{}, err
	}
	return &transactionMappings, nil
}
