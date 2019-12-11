package officialDocMethods

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"os"
)

//一开始自己网上搜的xorm的教程，不太完整，然后发现了官网文档，照着文档再来一遍，查漏补缺

//Engine和EngineGroup的区别，一个engine用于一个数据库。当有主从关系的多个数据库时可以选择EngineGroup.
//Engine
var engine *xorm.Engine
var engines []*xorm.Engine
var eg *xorm.EngineGroup

//新建一个engine(只管理一个数据库)
func CreateEngine() *xorm.Engine {
	engine, _ := xorm.NewEngine("mysql", "root:admin@tcp(127.0.0.1)/test?charset=utf8")
	return engine
}

//新建多个engine
func CreateEngines() []*xorm.Engine {
	for i:=0; i < 10; i++ {
		engine, _:= xorm.NewEngine("sqlite3", fmt.Sprintf("./test%d.db", i))
		engines = append(engines, engine)
	}
	return engines
}

func InitEngine() {
	if engine == nil {
		engine = CreateEngine()
	}
}

//Logs
//将日志记录在文件里面
func LogToFile() {
	f, err := os.Create("sql.log")
	if err != nil {
		println(err.Error())
		return
	}
	engine.SetLogger(xorm.NewSimpleLogger(f))
}

//syslog，不知道啥玩意。。。,貌似版本不对啥的，没有New这个方法...
//func Syslog() {
//	logWriter, err := syslog.New(syslog.LOG_DEBUG, "rest-xorm-example")
//	if err != nil {
//		log.Fatalf("Fail to create xorm system logger: %v\n", err)
//	}
//
//	logger := xorm.NewSimpleLogger(logWriter)
//	logger.ShowSQL(true)
//	engine.SetLogger(logger)
//}

//Connections pool连接池
func ConectionPool() {
	engine.SetMaxIdleConns(100)	//最大闲时连接数
	engine.SetMaxOpenConns(100)	//最大连接数
	engine.SetConnMaxLifetime(12000)		//最大存活时间
}

//Engine Group
func CreateEngineGroup1() *xorm.EngineGroup{
	conns := []string{
		"postgres://postgres:root@localhost:5432/test?sslmode=disable;", // first one is master
		"postgres://postgres:root@localhost:5432/test1?sslmode=disable;", // slave
		"postgres://postgres:root@localhost:5432/test2?sslmode=disable", // slave
	}

	eg, _:= xorm.NewEngineGroup("postgres", conns)
	return eg
}

func CreateEngineGroup2() *xorm.EngineGroup {
	var err error
	master, err := xorm.NewEngine("postgres", "postgres://postgres:root@localhost:5432/test?sslmode=disable")
	if err != nil {
		return eg
	}

	slave1, err := xorm.NewEngine("postgres", "postgres://postgres:root@localhost:5432/test1?sslmode=disable")
	if err != nil {
		return eg
	}

	slave2, err := xorm.NewEngine("postgres", "postgres://postgres:root@localhost:5432/test2?sslmode=disable")
	if err != nil {
		return eg
	}

	slaves := []*xorm.Engine{slave1, slave2}
	eg, err = xorm.NewEngineGroup(master, slaves)

	return eg
}

//Load balance policy
func LoadBalancePolicy() {
	conns := []string{
		"postgres://postgres:root@localhost:5432/test?sslmode=disable;",
		"postgres://postgres:root@localhost:5432/test1?sslmode=disable;",
		"postgres://postgres:root@localhost:5432/test2?sslmode=disable",
	}

	//1.RandomPolicy
	eg, _ = xorm.NewEngineGroup("postgres", conns, xorm.RandomPolicy())
	//2.WeightRandomPolicy
	//此时设置的test1数据库和test2数据库的随机访问权重为2和3
	eg, _ = xorm.NewEngineGroup("postgres", conns, xorm.WeightRandomPolicy([]int{2, 3}))
	//3.RoundRobinPolicy
	//轮询调度算法的原理是每一次把来自用户的请求轮流分配给内部中的服务器，从1开始，直到N(内部服务器个数)，然后重新开始循环。
	// 算法的优点是其简洁性，它无需记录当前所有连接的状态，所以它是一种无状态调度。
	eg, _ = xorm.NewEngineGroup("postgres", conns, xorm.RoundRobinPolicy())
	//4.WeightRoundRobinPolicy
	//此时设置的test1数据库和test2数据库的轮询访问权重为2和3
	eg, _ = xorm.NewEngineGroup("postgres", conns, xorm.WeightRoundRobinPolicy([]int{2, 3}))
	//5.LeastConnPolicy
	eg, _ = xorm.NewEngineGroup("postgres", conns, xorm.LeastConnPolicy())
	//6.Customerize Policy
	type GroupPolicy interface {
		Slave(group *xorm.EngineGroup) *xorm.Engine
	}
}
