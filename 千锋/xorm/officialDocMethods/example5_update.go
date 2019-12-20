package officialDocMethods

//6.Update
func UpdateEasy() {
	user := new(User)
	user.Name = "myname"
	engine.Id(1).Update(user)
}

//如果你想把0或者null更新到数据库（即某一个字段没有设置值，想用默认值更新它而不是）
func UpdateZeroOrEmpty() {
	user := new(User)
	engine.Id(1).Cols("age").Update(&user)
	engine.Id(1).AllCols().Update(&user)
	engine.Table(new(User)).Id(1).Update(map[string]interface{}{"age": 0})
	engine.Table("user").Id(1).Update(map[string]interface{}{"age": 0})
}

//6.1.Optimistic Lock 乐观锁
type User2 struct {
	Id int64
	Name string
	Version int `xorm:"version"`
}
func OptimisticLock(){
	var user User
	engine.Id(1).Get(&user)
	engine.Id(1).Update(&user)
	//读取出数据时，将此版本号一同读出，之后更新时，对此版本号加一。此时，将提交数据的版本数据与数据库表对应记录的当前版本信息进行
	//比对，如果提交的数据版本号等于数据库表当前版本号，则予以更新，否则认为是过期数据。

}

//updated标签跟created一样，会自动加上更新时间，为对应的当前系统时间