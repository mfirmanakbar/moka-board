package models

import (
	"time"

	"github.com/mfirmanakbar/moka-board/db"
	"github.com/mfirmanakbar/moka-board/objson"
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
	NotesJson      objson.LogNotes `gorm:"-"`
}

func (l *LoggingSummary) FindOneByConnectionId(ConnectionId int64, PageInput int) (*[]LoggingSummary, *utils.PaginatorResult, error) {
	var loggingsummries []LoggingSummary
	pgn := &utils.PaginatorResult{}

	dbres := db.JurnalMokaGorm.Where("connection_id = ?", ConnectionId)
	err := dbres.Error
	if err != nil {
		return &loggingsummries, pgn, err
	}

	pgn = utils.PaginatorUtil().Paging(&utils.PaginatorParam{
		DB:      dbres,
		Page:    PageInput,
		Limit:   10,
		OrderBy: []string{"id desc"},
	}, &loggingsummries)

	// for _, lss := range loggingsummries {
	// 	pingData := lss.Notes
	// 	pingJSON := make(map[string][]objson.LogNotes)
	// 	println(pingData)
	// 	json.Unmarshal([]byte(pingData), &pingJSON)
	// 	println(pingJSON)
	// 	// err = json.Unmarshal([]byte(pingData), &pingJSON)
	// 	// if err != nil {
	// 	// 	println(err.Error())
	// 	// }
	// }

	return &loggingsummries, pgn, nil
}

type LoggingSummaryInterface interface {
	FindOneByConnectionId(ConnectionId int64, Page int) (*[]LoggingSummary, *utils.PaginatorResult, error)
}

func LoggingSummaryDTO() LoggingSummaryInterface {
	return &LoggingSummary{}
}
