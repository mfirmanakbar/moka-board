package db

import (
	"os"
)

var (
	DatabaseDriver   = "mysql"
	DatabaseUser     = os.Getenv("JURNAL_MOKA_DB_USERNAME")
	DatabasePassword = os.Getenv("JURNAL_MOKA_DB_PASSWORD")
	DatabasePort     = os.Getenv("JURNAL_MOKA_DB_PORT")
	DatabaseHost     = os.Getenv("JURNAL_MOKA_DB_HOST")
	DatabaseName     = os.Getenv("JURNAL_MOKA_DB_NAME")
)
