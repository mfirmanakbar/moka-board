package models

import (
	"github.com/mfirmanakbar/moka-board/datasources"
	"time"
)

type Connection struct {
	Id                         int64     `json:"id"`
	AccountMappingId           int64     `json:"account_mapping_id"`
	Name                       string    `json:"name"`
	OutletName                 string    `json:"outlet_name"`
	OutletId                   int64     `json:"outlet_id"`
	WarehouseId                int64     `json:"warehouse_id"`
	TagIds                     string    `json:"tag_ids"`
	PaymentMethod              string    `json:"payment_method"`
	PaymentMethodV2            string    `json:"payment_method_v_2"`
	ProductBuyAccountId        int64     `json:"product_buy_account_id"`
	ProductSellAccountId       int64     `json:"product_sell_account_id"`
	DefaultCustomerId          int64     `json:"default_customer_id"`
	RoundingAccountId          int64     `json:"rounding_account_id"`
	RoundingProductId          int64     `json:"rounding_product_id"`
	Prefix                     string    `json:"prefix"`
	Suffix                     string    `json:"suffix"`
	StartDate                  time.Time `json:"start_date"`
	IsTrackInventory           int       `json:"is_track_inventory"`
	TrackInventoryAccountId    int64     `json:"track_inventory_account_id"`
	LastUpdatedSyncMaster      time.Time `json:"last_updated_sync_master"`
	LastSuccessSyncMaster      time.Time `json:"last_success_sync_master"`
	LastUpdatedSyncTransaction time.Time `json:"last_updated_sync_transaction"`
	LastSuccessSyncTransaction time.Time `json:"last_success_sync_transaction"`
	CreatedAt                  time.Time `json:"created_at"`
	UpdatedAt                  time.Time `json:"updated_at"`
	DeletedAt                  time.Time `json:"deleted_at"`
	SyncStatusMaster           int       `json:"sync_status_master"`
	SyncStatusTransaction      int       `json:"sync_status_transaction"`
	DailySyncTime              string    `json:"daily_sync_time"`
	IsDailySync                int       `json:"is_daily_sync"`
	SyncTrnByProductName       int       `json:"sync_trn_by_product_name"`
	AddTrxNoIdentifier         int       `json:"add_trx_no_identifier"`
}

type ConnectionParams struct {
	CompanyId       int64
	ConnectionId    int64
	ConnectionName  string
	ShowDeleted     bool
	ShowOnlySyncing bool
}

type ConnectionInterface interface {
	SearchData(ConnectionParams) (*[]Connection, error)
	QueryParams(ConnectionParams) map[string]interface{}
}

func ConnectionDTO() ConnectionInterface {
	return &Connection{}
}

func (c Connection) SearchData(prm ConnectionParams) (*[]Connection, error) {
	var err error
	var connections []Connection

	if prm.CompanyId < 1 && prm.ConnectionId < 1 && !prm.ShowOnlySyncing {
		return &[]Connection{}, err
	}

	err = datasources.JmDb.Where(c.QueryParams(prm)).Unscoped().Find(&connections).Error
	if err != nil {
		return &connections, err
	}
	return &connections, nil
}

func (c Connection) QueryParams(prm ConnectionParams) map[string]interface{} {
	queryParams := make(map[string]interface{})
	if prm.ConnectionId > 0 {
		queryParams["id"] = prm.ConnectionId
	}
	if !prm.ShowDeleted {
		queryParams["deleted_at"] = "IS NULL"
	}
	return queryParams
}
