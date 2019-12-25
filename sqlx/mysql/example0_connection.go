package mysqlx

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

//func Conct()(){
//		driverName := "mysql"
//		dataSourceName := "root:admin@tcp(127.0.0.1:3306)/test?charset=utf8"
//		var err error
//		Db, err =sqlx.Connect(driverName, dataSourceName)
//		if err != nil{
//			fmt.Printf("Connect to database successful!")
//		}
//}

func init(){
	var err error
	driverName := "mysql"
	dataSourceName := "root:admin@tcp(127.0.0.1:3306)/test?charset=utf8"
	Db, err = sqlx.Connect(driverName, dataSourceName)
	if err != nil{
		fmt.Printf("%s", err)
	}
}