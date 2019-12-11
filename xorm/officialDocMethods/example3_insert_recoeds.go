package officialDocMethods

import (
	"fmt"
	"time"
)

type Question struct {
	id int64
	Content string
}

//Insert data
//插入一条数据
func InsertOneRecord() {
	user := new(User)
	user.Name = "myname"
	_, _ = engine.Insert(user)

	fmt.Println(user.Id)
}

//通过切片插入多条数据
func InsertMultipleRecordsBySlice() {
	users := make([]User, 1)
	users[0].Name = "name0"

	_,_ := engine.Insert(&users)
}

//通过切片指针插入多条数据
func InsertMultipleRecordsBySliceofPointer() {
	users := make([]*User, 1)
	users[0] = new(User)
	users[0].Name = "name0"
	_,_ := engine.Insert(&users)
}

//一次插入到两个表中
func InsertOneRecordOnTwoTable() {
	user := new(User)
	user.Name = "myname"
	question := new(Question)
	question.Content = "whywhywhy?"
	_,_ :=engine.Insert(user, question)
}

//多条数据插入到多个表中
func InsertMultipleRecordsOnMultipleTables() {
	users := make([]User, 1)
	users[0].Name = "name0"
	questions := make([]Question, 1)
	questions[0].Content = "whywhywhy?"
	_,_ :=engine.Insert(&users, &questions)
}

//插入一条或多条记录到多个表格中
func InsertOneOrMultipleRecordsOnMultipleTables() {
	user := new(User)
	user.Name = "myname"
	questions := make([]Question, 1)
	questions[0].Content = "whywhywhy?"
	_,_ := engine.Insert(user, &questions)
}

//如果想用事务的话，需要在insert之前用sessoin.begin

//Created
type UserInfo1 struct {
	Id int64
	Name string
	CreatedAt time.Time `xorm: "created"`	//这个标签会记录时间
}

type JsonTime time.Time
func (j JsonTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(j).Format("2006-01-02 15:04:05")+`"`), nil
}

type UserInfo2 struct {
	Id int64
	Name string
	CreatedAt JsonTime `xorm:"created"`
}

type UserInfo3 struct {
	Id int64
	Name string
	CreatedAt int64 `xorm:"created"`
}

//一般engine会用本地时区，如果想修改时区。使用 engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai)
//当调用Insert(),InsertOne(), user.CreatedAt会自动被time.now()或者time.Now().Unix()填充