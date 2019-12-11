package officialDocMethods

import "time"

func SQLQuery() {
	sql := "select * from user"
	_, _ = engine.Query(sql)
}

func SQLCommand() {
	sql := "update user set name = ? where id = ?"
	_, _ = engine.Exec(sql, "xiaolun", 1)
}

//事务transaction
func TransactionTest() {
	session := engine.NewSession()
	defer session.Close()

	// add Begin() before any action
	err := session.Begin()
	user1 := User{Name: "xiaoxiao", CreatedAt: time.Now()}
	_, err = session.Insert(&user1)
	if err != nil {
		session.Rollback()
		return
	}
	user2 := User{Name: " yyy"}
	_, err = session.Where("id = ?", 2).Update(&user2)
	if err != nil {
		session.Rollback()
		return
	}

	_, err = session.Exec("delete from user where name = ?", user2.Name)
	if err != nil {
		session.Rollback()
		return
	}

	// add Commit() after all actions
	err = session.Commit()
	if err != nil {
		return
	}
}
//这个方法表达的意思就是上述操作要么全部成功要么全部失败（有一个失败就算全部失败），全部成功则提交更新到数据库，
// 有一个失败或多个失败则回滚到操作前。
//这多个操作构成一个事务，也就是begin()和commit()之间的操作。