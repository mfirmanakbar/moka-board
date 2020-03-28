package datasources

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql database driver
	"log"
)

type JurnalMokaDbGorm struct {
	DB *gorm.DB
}

//func (jurnalMokaDbGorm *JurnalMokaDbGorm) Initialize(DbDriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
//	var err error
//
//	DbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
//	jurnalMokaDbGorm.DB, err = gorm.Open(DbDriver, DbUrl)
//	if err != nil {
//		fmt.Printf("Cannot connect to %s database", DbDriver)
//		log.Fatal("This is the error:", err)
//	} else {
//		fmt.Printf("We are connected to the %s database", DbDriver)
//	}
//	fmt.Println("")
//}

func Initialize(DbDriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*gorm.DB, error) {
	var err error
	var conn *gorm.DB

	DbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	conn, err = gorm.Open(DbDriver, DbUrl)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", DbDriver)
		log.Fatal("This is the error:", err)
		fmt.Println("")
		return nil, err
	} else {
		fmt.Printf("We are connected to the %s database", DbDriver)
		fmt.Println("")
		return conn, nil
	}
}
