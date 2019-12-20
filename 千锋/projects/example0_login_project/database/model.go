package database

import "fmt"

var (
	createSQLs = []string{
		"CREATE TABLE IF NOT EXISTS `user` (\n" +
			" `id` int(32) NOT NULL AUTO_INCREMENT, \n" +
			" `user_name` varchar(32) NOT NULL, \n" +
			" `password` varchar(32) NOT NULL, \n" +
			" `is_admin` bool DEFAULT false, \n" +
			" `admin_id` int(32) DEFAULT NULL, \n" +
			" PRIMARY KEY (`id`)\n" +
			" ) ENGINE=InnoDB DEFAULT CHARSET=utf8;",
		//"CREATE TABLE IF NOT EXISTS `city` (\n" +
		//	" `id` int(32) NOT NULL AUTO_INCREMENT, \n" +
		//	" `name` varchar(32) DEFAULT NULL\n" +
		//	" ) ENGINE=InnoDB DEFAULT CHARSET=utf8",
		//"CREATE TABLE IF NOT EXISTS `admin` (\n" +
		//	" `admin_id`"
	} //建表语句之类
)

//type Admin struct {
	//如果field名称为Id，而且类型为int64，并没有定义tag，则会被xorm视为主键，并且拥有自增属性
	//AdminId    int64     `xorm:"'admin_id' pk autoincr" json:"id"`
	//AdminName  string    `xorm:"'admin_name' varchar(32)" json:"admin_name"`
	//CreateTime time.Time `xorm:"'create_time' DateTime" json:"create_time"`
	//Status     int64     `xorm:"'status' default 0" json:"status"`
	//Avatar     string    `xorm:"'avatar' varchar(255)" json:"avatar"`
	//Pwd        string    `xorm:"'pwd' varchar(255)" json:"pwd"`      //管理员密码
	//CityName   string    `xorm:"'city_name' varchar(12)" json:"city_name"` //管理员所在城市名称
	//CityId     int64     `xorm:"index" json:"city_id"`
	//City       *City     `xorm:"- <- ->"` //所对应的城市结构体（基础表结构体）
//}

//type City struct {
//	Id   int64
//	Name string
//}

type User struct {
	Id       int64  `xorm:"pk autoincr" json:"id"`
	UserName string `xorm:"user_name varchar(32)" json:"user_name"`
	Password string `xorm:"password varchar(32)" json:"password"`
	IsAdmin	bool `xorm:"is_admin bool" json:"is_admin"`
	AdminId int64 `xorm:"admin_id int(32)" json:"admin_id"`
}

func (u *User) String() string {
	return fmt.Sprintf("Id: %d,\n Username: %s, \n Password: %s, \n IsAdmin: %d, \n AdminId: %d, \n", u.Id,
		u.UserName, u.Password, u.IsAdmin, u.AdminId)
}



