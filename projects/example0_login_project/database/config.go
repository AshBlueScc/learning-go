package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	utils2 "xorm/projects/example0_login_project/utils"
)

var (
	dbMySQLPoolSize = utils2.GetEnvInt("MYSQL_POOL_SIZE", 50)
	dbMySQLUser     = utils2.GetEnvStr("MYSQL_USER", "root")
	dbMySQLPassword = utils2.GetEnvStr("MYSQL_PASSWORD", "admin")
	dbMySQLHost     = utils2.GetEnvStr("MYSQL_HOST", "localhost")
	dbMySQLPort     = utils2.GetEnvInt("MYSQL_PORT", 3306)
	dbMySQLDB       = utils2.GetEnvStr("MYSQL_DB", "test")

	dbTimeout = utils2.GetEnvInt("DB_TIMEOUT", 10)
	dbURL     = utils2.GetEnvStr("DB_URL", "root:admin@tcp(127.0.0.1)/test?charset=utf8")
	db        *xorm.Engine
)

func init() {
	if len(dbURL) <= 0 {
		dbURL = fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=true",
			dbMySQLUser, dbMySQLPassword, dbMySQLHost, dbMySQLPort, dbMySQLDB)
	}
	var err error
	db, err = xorm.NewEngine("mysql", dbURL)
	db.ShowSQL(true)
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(dbMySQLPoolSize / 2)
	db.SetMaxOpenConns(dbMySQLPoolSize)

	for _, sql := range createSQLs {
		_, err = db.Exec(sql)
		if err != nil {
			panic(err)
		}
	}
}

func GetDB() *xorm.Engine {
	return db
}

