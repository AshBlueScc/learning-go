package officialDocMethods

import "time"

//1.xorm 标签和json怎么同时使用？
type User3 struct {
	Name string `json: "name" xorm:"name"`
}

//2.xorm支持l联合主键嘛？
type User4 struct {
	id int64 `id(xorm.PK{1, 2})`
	Name string
}
//第一个域和第二个域作为联合主键

//3.怎么使用join?
type Userinfo struct {
	Id int64
	Name string
	DetailId int64
}

type Userdetail struct {
	Id int64
	Gender int
}

type User5 struct {
	Userinfo `xorm:"extends"`
	Userdetail `xorm:"extends"`
}

func JoinTest() {
	var users = make([]User5, 0)
	_ = engine.Table(&Userinfo{}).Join("LEFT", "userdetail", "userinfo.detail_id = userdetail.id").Find(&users)
	_ = engine.SQL("select * from userinfo, userdetail where userinfo.detail_id = userdetail.id").Find(&users)

	//assert(User.Userinfo.Id != 0 && User.Userdetail.Id != 0)
}

//Please notice that Userinfo field on User should be before Userdetail because of the order on join SQL stsatement.
// If the order is wrong, the same name field may be set a wrong value and no error will be indicated.

//4.怎么设置数据库时区?
func SetTimeZone() {
	location, _ := time.LoadLocation("Asia/Shanghai")
	engine.TZLocation = location
}