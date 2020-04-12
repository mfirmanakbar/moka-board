package models

import "time"

type LoggingDetail struct {
	ID             int64      `json:"id"`
	SummaryId      int64      `json:"summary_id"`
	LoggingType    int        `json:"logging_type"`
	Notes          string     `json:"notes"`
	Payload        string     `json:"payload"`
	SyncStatus     int        `json:"sync_status"`
	CreatedAt      *time.Time `sql:"index" json:"created_at"`
	UpdatedAt      *time.Time `sql:"index" json:"updated_at"`
	DeletedAt      *time.Time `sql:"index" json:"deleted_at"`
	MarkAsResolved int        `json:"mark_as_resolved"`
}
