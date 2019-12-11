package mysql

func Insert(user *User) (int64, bool) {
	engine := GetEngine()
	engine.ShowSQL(true)
	affected, err := engine.Insert(user)
	if err != nil {
		return affected, false
	}
	return affected, true
}
