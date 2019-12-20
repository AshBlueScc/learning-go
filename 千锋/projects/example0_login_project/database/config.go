package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm/千锋/projects/example0_login_project/utils"
)

var (
	dbMySQLPoolSize = utils.GetEnvInt("MYSQL_POOL_SIZE", 50)
	dbMySQLUser     = utils.GetEnvStr("MYSQL_USER", "root")
	dbMySQLPassword = utils.GetEnvStr("MYSQL_PASSWORD", "admin")
	dbMySQLHost     = utils.GetEnvStr("MYSQL_HOST", "localhost")
	dbMySQLPort     = utils.GetEnvInt("MYSQL_PORT", 3306)
	dbMySQLDB       = utils.GetEnvStr("MYSQL_DB", "test")

	dbTimeout = utils.GetEnvInt("DB_TIMEOUT", 10)
	dbURL     = utils.GetEnvStr("DB_URL", "root:admin@tcp(127.0.0.1)/test?charset=utf8")
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

