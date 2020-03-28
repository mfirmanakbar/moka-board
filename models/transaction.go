package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type TransactionMappings struct {
	Id                  int64
	JurnalTransactionId int64
	MokaTransactionId   int64
	MokaTransactionType int8
	ConnectionId        int64
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           time.Time
}

type TransactionMappingList []*TransactionMappings

func init() {
	orm.RegisterModel(new(TransactionMappings))
}
