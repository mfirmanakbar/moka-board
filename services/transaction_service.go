package services

//
//import (
//	"fmt"
//	"github.com/astaxie/beego/orm"
//	"github.com/mfirmanakbar/moka-board/models"
//)
//
//func FindTransactionMappings(
//	connectionId int64,
//	mokaTransactionId int64,
//	jurnalTransactionId int64,
//	mokaTransactionType int64,
//) {
//	o := orm.NewOrm()
//	var tm models.TransactionMappingList
//
//	qs := o.QueryTable("transaction_mappings")
//	//if connectionId > 0 {
//	//	qs = qs.Filter("connection_id", connectionId)
//	//}
//	//if mokaTransactionId > 0 {
//	//	qs = qs.Filter("moka_transaction_id", mokaTransactionId)
//	//}
//	//if jurnalTransactionId > 0 {
//	//	qs = qs.Filter("jurnal_transaction_id", jurnalTransactionId)
//	//}
//
//	//if mokaTransactionType > -1 {
//	//	qs = qs.Filter("moka_transaction_type", mokaTransactionType)
//	//}
//
//	qs = FilterQsMoreThan(qs, "connection_id", connectionId, 0)
//	qs = FilterQsMoreThan(qs, "moka_transaction_id", mokaTransactionId, 0)
//	qs = FilterQsMoreThan(qs, "jurnal_transaction_id", jurnalTransactionId, 0)
//	qs = FilterQsMoreThan(qs, "moka_transaction_type", mokaTransactionType, -1)
//
//	_, err := qs.All(&tm)
//
//	if err == orm.ErrNoRows {
//		fmt.Println("No result found.")
//	} else if err == orm.ErrMissPK {
//		fmt.Println("No primary key found.")
//	} else {
//		//fmt.Printf("Returned Rows Num: %s, %s", num, err)
//		//fmt.Println("")
//		for i, s := range tm {
//			fmt.Println(i, s)
//		}
//	}
//}
//
//func FilterQsMoreThan(qs orm.QuerySeter, columnName string, param int64, diff int64) orm.QuerySeter {
//	if param > diff {
//		qs = qs.Filter(columnName, param)
//		return qs
//	}
//	return qs
//}
