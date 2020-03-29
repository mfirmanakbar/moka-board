package datasources

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

var (
	JmDb       *gorm.DB
	DbDriver   = "mysql"
	DbUser     = os.Getenv("JURNAL_MOKA_DB_USERNAME")
	DbPassword = os.Getenv("JURNAL_MOKA_DB_PASSWORD")
	DbPort     = os.Getenv("JURNAL_MOKA_DB_PORT")
	DbHost     = os.Getenv("JURNAL_MOKA_DB_HOST")
	DbName     = os.Getenv("JURNAL_MOKA_DB_NAME")
)

func init() {
	var err error
	DbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	JmDb, err = gorm.Open(DbDriver, DbUrl)
	JmDb.LogMode(true)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", DbDriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", DbDriver)
	}
	fmt.Println("")
}
