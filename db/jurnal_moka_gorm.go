package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // lib for mysql void type
	"github.com/mfirmanakbar/moka-board/myenv"
)

// JurnalMokaGorm is out DB config
var (
	JurnalMokaGorm *gorm.DB
)

func init() {
	var err error
	DbURL := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		myenv.DatabaseUser,
		myenv.DatabasePassword,
		myenv.DatabaseHost,
		myenv.DatabasePort,
		myenv.DatabaseName,
	)
	JurnalMokaGorm, err = gorm.Open(myenv.DatabaseDriver, DbURL)
	JurnalMokaGorm.LogMode(true)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", myenv.DatabaseDriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", myenv.DatabaseDriver)
	}
	fmt.Println("")
}
