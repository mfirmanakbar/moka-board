package myenv

import "os"

var (
	DatabaseDriver   = "mysql"
	DatabaseUser     = os.Getenv("JURNAL_MOKA_DB_USERNAME")
	DatabasePassword = os.Getenv("JURNAL_MOKA_DB_PASSWORD")
	DatabasePort     = os.Getenv("JURNAL_MOKA_DB_PORT")
	DatabaseHost     = os.Getenv("JURNAL_MOKA_DB_HOST")
	DatabaseName     = os.Getenv("JURNAL_MOKA_DB_NAME")
	BaseURLOldFe     = os.Getenv("BASE_URL_OLD_FE")
	BaseURLNewFe     = os.Getenv("BASE_URL_NEW_FE")
)
