package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

var Engine *xorm.Engine

//为了复用，发现写在main外面更方便点...
func connect() *xorm.Engine {
	driverName := "mysql"
	dataSourceName := "root:admin@tcp(127.0.0.1:3306)/test?charset=utf8"
	engine, err := xorm.NewEngine(driverName, dataSourceName)

	if err != nil {
		panic("访问数据库发生错误！")
	}

	return engine
}

//单例模式
func GetEngine() *xorm.Engine {
	if Engine == nil {
		Engine = connect()
	}
	return Engine
}

func ConnectionTest() {
	engine := GetEngine()
	//	//1.修改表名方法一
	//	//单独设置表名的映射
	//	//engine.SetTableMapper(core.SameMapper{})
	//	//单数设置字段名的映射
	//	//engine.SetColumnMapper(core.GonicMapper{})
	//
	//
	//	//2.修改表名方法二
	//	//给每个表加前缀包括加后缀：reference:https://www.2cto.com/kf/201905/809852.html
	//	//给表名category加前缀t_查询
	engine.SetTableMapper(core.NewPrefixMapper(core.SnakeMapper{}, "t_"))
	//	//给查询的表名加后缀_suffix，在结构体按照驼峰转小写（SnakeMapper）的基础上
	//	//engine.SetTableMapper(core.NewSuffixMapper(core.SnakeMapper{}, "_suffix"))

	//这个修改表名跟方法一一样
	//结构体名称到表名和结构体field到表字段的映射
	//这个放在setTableMapper后面会覆盖setTableMapper，放在前面则不会。这个相当于是对表和字段各使用一次，所以要么分开写，要么就用这一个
	//engine.SetMapper(core.SnakeMapper{})

	//是否显示SQL语句（开发调试时使用）
	engine.ShowSQL(true)
	//设置数据库最大连接数
	engine.SetMaxOpenConns(10)
	//设置最大空闲连接数量：默认是2
	engine.SetMaxIdleConns(5)

	var category = new(Category)

	//想通过id来查询，结构体里面必须加tag将其描述为pk否则不可用
	_, err := engine.Id(3).Get(category)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(category)

	//4.修改表名方法四
	//利用Table(表名)这个方法
	//_, err = engine.Table("t_category").Id(3).Get(category)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Print(category)

}

//3.修改表名方法三
//让结构体Category实现TableName这个方法，修改表名
//func (category *Category) TableName() string {
//	return "t_category"
//}
