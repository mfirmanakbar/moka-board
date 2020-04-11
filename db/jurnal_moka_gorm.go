package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var (
	JurnalMokaGorm *gorm.DB
)

func init() {
	var err error
	DbUrl := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DatabaseUser,
		DatabasePassword,
		DatabaseHost,
		DatabasePort,
		DatabaseName,
	)
	JurnalMokaGorm, err = gorm.Open(DatabaseDriver, DbUrl)
	JurnalMokaGorm.LogMode(true)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", DatabaseDriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", DatabaseDriver)
	}
	fmt.Println("")
}
