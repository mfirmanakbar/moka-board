package db

const (
	SelectAllColumns = `SELECT c.id, c.account_mapping_id, c.name, c.outlet_name, c.outlet_id, c.warehouse_id, c.tag_ids, 
		c.payment_method, c.payment_method_v2, c.product_buy_account_id, c.product_sell_account_id, c.default_customer_id, 
		c.rounding_account_id, c.rounding_product_id, c.prefix, c.suffix, c.start_date, c.is_track_inventory, 
		c.track_inventory_account_id, c.last_updated_sync_master, c.last_success_sync_master, c.last_updated_sync_transaction, 
		c.last_success_sync_transaction, c.created_at, c.updated_at, c.deleted_at, c.sync_status_master, c.sync_status_transaction, 
		c.daily_sync_time, c.is_daily_sync, c.sync_trn_by_product_name, c.add_trx_no_identifier 
		FROM connections c `
)
