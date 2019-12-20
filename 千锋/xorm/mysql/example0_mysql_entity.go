package mysql

type Category struct {
	Id           int `xorm:"pk autoincr"`
	CategoryId   int
	CategoryName string
	Logo         string
}

type User struct {
	ID   int    `xorm:"id pk autoincr"`
	Name string `xorm:"name"`
	Age  int    `xorm:"age"`
}

func (user *User) TableName() string {
	return "prefix_user"
}

