package mysql

import (
	"log"
	"xorm.io/core"
)

type User struct {
	ID   int    `xorm:"id"`
	Name string `xorm:"name"`
	Age  int    `xorm:"age"`
}

func CreateTablesTest() {
	var engine = GetEngine()

	//给表名加前缀prefix_
	engine.SetTableMapper(core.NewPrefixMapper(core.GonicMapper{}, "prefix_"))
	//执行之后，结构体 `type User struct` 默认对应的表名就变成了 `prefix_user` 了
	//而之前默认的是 `user`
	err := engine.CreateTables(User{})
	if err != nil {
		log.Fatal(err)
		return
	}

	//给表名加后缀_suffix
	engine.SetTableMapper(core.NewSuffixMapper(core.GonicMapper{}, "_suffix"))
	//执行之后，结构体 `type User struct` 默认对应的表名就变成了 `prefix_user` 了
	//而之前默认的是 `user`
	err = engine.CreateTables(User{})
	if err != nil {
		log.Fatal(err)
		return
	}

}
