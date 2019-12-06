package mysql

import (
	"github.com/go-xorm/xorm"
)

func InsertTableTest() {
	user := new(User)
	user.Name = "testInsert"
	user.Age = 18

	insert(user, GetEngine())
}

func insert(user *User, engine *xorm.Engine) (int64, bool) {
	engine.ShowSQL(true)
	affected, err := engine.Insert(user)
	if err != nil {
		return affected, false
	}
	return affected, true
}
