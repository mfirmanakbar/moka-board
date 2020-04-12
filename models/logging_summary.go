package models

import (
	"time"

	"github.com/mfirmanakbar/moka-board/db"
	utils "github.com/mfirmanakbar/moka-board/utils/paginator_util"
)

type LoggingSummary struct {
	ID             int64           `json:"id"`
	ConnectionId   int64           `json:"connection_id"`
	SyncDate       *time.Time      `sql:"index" json:"sync_date"`
	StatusType     int             `json:"status_type"`
	MethodSyncType int             `json:"method_sync_type"`
	DataType       int             `json:"data_type"`
	Notes          string          `json:"notes"`
	CreatedAt      *time.Time      `sql:"index" json:"created_at"`
	UpdatedAt      *time.Time      `sql:"index" json:"updated_at"`
	DeletedAt      *time.Time      `sql:"index" json:"deleted_at"`
	LoggingDetails []LoggingDetail `gorm:"foreignkey:SummaryId"`
}

func (l *LoggingSummary) FindOneByConnectionId(id int64) (*[]LoggingSummary, *utils.PaginatorResult, error) {
	var loggingsummries []LoggingSummary
	pgn := &utils.PaginatorResult{}

	dbres := db.JurnalMokaGorm.Where("connection_id = ?", id)
	err := dbres.Error
	if err != nil {
		return &loggingsummries, pgn, err
	}

	pgn = utils.PaginatorUtil().Paging(&utils.PaginatorParam{
		DB:      dbres,
		Page:    1,
		Limit:   10,
		OrderBy: []string{"id desc"},
	}, &loggingsummries)

	return &loggingsummries, pgn, nil
}

type LoggingSummaryInterface interface {
	FindOneByConnectionId(id int64) (*[]LoggingSummary, *utils.PaginatorResult, error)
}

func LoggingSummaryDTO() LoggingSummaryInterface {
	return &LoggingSummary{}
}