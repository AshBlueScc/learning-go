package officialDocMethods

import (
	"fmt"
)

//5. Chainable APIs
//5.1. Chainable APIs for Queries, Execusions and Aggregations
func ChainableAPIsTest() {
	engine.Alias("tableName").Where("tableName.columnName = ?", "columnName").Get(new(User))
	engine.Select("a.*, (select name from b limit 1) as name").Find(new(User))
	////连接词
	//And,Or,In(可以用builder.Builder作为一个子查询)
	////操作词
	//Asc,Desc,OrderBy,Select,SQL,Where,Omit,Distinct,Table,Limit,Join,GroupBy,Having
	////指定字段
	//ID，如果主键是联合主键可以用ID(core.PK{1, "name"})，Cols,AllCols,MustCols
}

//5.2. Override default behavior APIs
//NoAutoTime()		No auto timestamp for Created and Updated fields for INSERT and UPDATE
//NoCache()			Disable cache lookup
//NoAutoCondition()	Disable auto generate condition from bean. For example:
//	eg.
//	engine.Where("name = ?", "lunny").Get(&User{Id:1})
//	SELECT * FROM user where name='lunny' AND id = 1 LIMIT 1
//	engine.Where("name = ?", "lunny").NoAutoCondition().Get(&User{Id:1})
//	SELECT * FROM user where name='lunny' LIMIT 1
//
//UseBool(…string)
//xorm’s default behavior is fields with 0, “”, nil, false, will not be used during query or update, use this method toexplicit specify bool type fields for query or update
//
//Cascade(bool)
//Do cascade lookup for associations

//5.3.Get one record
func GetOneRecord() {
	var user = User{Id:27}
	_, _ = engine.Get(&user)
	// or has, err := engine.Id(27).Get(&user)

	var user1 = User{Name:"xlw"}
	_, _ = engine.Get(&user1)
}

//5.4.Find
//Fetch multipe objects into a slice or a map, use Find：
//get和find的区别就是get一般都是取一个bean,find取多个beans
func GetMultipleRecords() {
	//Fetch multipe objects into a slice or a map, use Find：
	var everyone []User
	_ = engine.Find(&everyone)

	users := make(map[int64]User)
	_ = engine.Find(&users)

	//also you can use Where, Limit
	var allusers []User
	_ = engine.Where("id > ?", "3").Limit(10, 20).Find(&allusers)

	//or you can use a struct query
	var tenusers []User
	_ = engine.Limit(10).Find(&tenusers, &User{Name:"xlw"}) //Get All Name="xlw" limit 10 offset 0

	//or In function
	var threeusers []User
	_ = engine.In("id", 1, 3, 5).Find(&threeusers)	//Get All id in (1, 3, 5)

	//The default will query all columns of a table. Use Cols function if you want to select some columns
	var twocolumns []User
	_ = engine.Cols("id", "name").Find(&twocolumns)	//Find only id and name

	//You can also use slice of ints if you only want one column
	var ints []int64
	_ = engine.Table("user").Cols("id").Find(&ints)
}


//join usage
type Group struct {
	Id int64
	Name string
}

type User1 struct {
	Id int64
	Name string
	GroupId int64 `xorm:"index"`
}

//query all user and his group name:
type UserGroup struct {
	User1 `xorm:"extends"`
	Name string
}

func (UserGroup) TableName() string {
	return "user1"
}

func AllUsersAndGroupName() {
	var users []UserGroup
	_ = engine.Join("INNER", "group", "group.id = user1.group_id").Find(&users)

	var users1 []UserGroup
	_ = engine.SQL("select user.*, group.name from user, group where user.group_id = group.id").Find(&users1)
}

type Type struct {
	Id int64
	Name string
}

type UserGroupType struct {
	User `xorm:"extends"`
	Group `xorm:"extends"`
	Type `xorm:"extends"`
}

func ThreeTableJoins() {
	var users []UserGroupType
	_ = engine.Table("user1").Join("INNER", "group", "group.id = user1.group_id").
		Join("INNER", "type", "type.id = user.type_id").Find(&users)
}


//5.6.Iterate records
func IterateUsers() {
	_ = engine.Where("age > ? or name = ?", 30, "xlw").Iterate(new(User), func(idx int, bean interface{}) error {
		fmt.Println("%d : %#v", idx, bean.(*User))
		return nil
	})
}

//5.7.Count method usage
func CountUsers() {
	user := new(User)
	total, _ := engine.Where("id > ?", 1).Count(user)
	fmt.Print(total)
}

//Rows
func RowTest() {
	rows, _ := engine.Rows(new(User))
	defer rows.Close()

	bean := new(User)
	// rows.Next() 返回一个bool，是否有下一个数据
	for rows.Next(){
		// Scan 扫描记录
		if err := rows.Scan(bean);err!=nil {
			fmt.Print("err:", err)
		}else {
			fmt.Printf("%#v\n", bean)
		}
	}
}

type SumStruct struct {
	Id int64
	Money int
	Rate float32
}

//Sum
func SumTest() {
	total, _ := engine.Sum(new(SumStruct), "monry")
	fmt.Printf("money is %d", int(total))
}

//SumInt
func SumIntTest() {
	total, _ := engine.SumInt(new(SumStruct), "monry")
	fmt.Printf("money is %d", total)
}

//Sums
func SumsTest() {
	var totals []float64
	totals, _ = engine.Sums(new(SumStruct), "money", "rate")
	fmt.Printf("money is %d, rate is %.2f", totals[0], totals[1])
}

//SumsInt
func SumsIntTest() {
	var totals []int64
	totals, _ = engine.SumsInt(new(SumStruct), "money")
	fmt.Printf("money is %d", totals[0])
}