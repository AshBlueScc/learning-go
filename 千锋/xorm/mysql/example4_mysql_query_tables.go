package mysql

import (
	"log"
)



//支持sql查询
func Query() {
	engine := GetEngine()
	engine.ShowSQL(true)

	//1.查询一个string类型的sql,返回[]map[string][]byte  类型的切片(查询)
	//results是一个切片，key是从0开始的int,value是map.
	//map的key是表格的字段名，value是值
	results, err := engine.Query("select * from prefix_user")
	if err != nil {
		return
	}
	for _, maps := range results {
		for key, value := range maps {
			log.Println(key, string(value))
		}
	}

	//2、执行一个string的sql，返回结果影响行数（增删改）
	//rows为影响行数，insertId如果是插入新增了id则返回那个id否则返回0
	affected, err := engine.Exec("update prefix_user set age = 22 where id = 4")
	if err != nil {
		return
	}
	rows, _ :=affected.RowsAffected()
	insertId, _ := affected.LastInsertId()
	println(rows, insertId)
}

//ORM方法
func Query1() {
	engine := GetEngine()
	engine.ShowSQL(true)

	//主要有七个方法和其他辅助方法来操作数据库
	//1.插入一个或者多个数据:
	//
	//user := new(User)
	//user.Age = 1
	//user.Name = "one"
	//
	////原po主传指针我也照着来，结果不行，这里不能传指针，传值就行了
	//affected, err := engine.Insert(user)
	//if err != nil {
	//	return
	//}
	//println(affected)
	//
	//user1 := new(User)
	//user1.Age = 2
	//user1.Name = "two"
	//
	//user2 := new(User)
	//user2.Age = 3
	//user2.Name = "three"
	//affected1, err := engine.Insert(user1, user2)
	//if err != nil {
	//	return
	//}
	//println(affected1)

	////2.从数据库里面查询一条记录
	//var user21 = new(User)
	////可以添加id就按照id查询，不加则是limit第一个
	////user21.ID = 3
	////不能传指针进去，可以这种形式&User{ID: 3}, 但是无法取到查询的对象的值
	//has, err := engine.Get(user21)
	//if err != nil {
	//	print(has)
	//	return
	//}
	//if has {
	//	println(user21.ID, user21.Name, user21.Age)
	//}

	////3.从数据库中查询多条记录
	//var everyone []User
	//err := engine.Find(&everyone)
	//if err != nil {
	//	return
	//}
	//for _, user := range everyone {
	//	println(user.Name)
	//}

	////4.查询多条记录，然后每条进行处理，有两个方法，一个是iterator,另一个是raw：
	////Iterate的扩展
	////engine.Cols("name").Iterate(new(Account), echo)  // 查询特定字段
	////engine.Omit("name").Iterate(new(Account), echo)  // 排除特定字段
	////engine.Limit(3,2).Iterate(new(Account), echo)        // 查询结果偏移
	//
	//usr := new(User)
	//err := engine.Iterate(new(User), func(idx int, bean interface{}) error {
	//	fmt.Printf("%d : %#v \n", idx, bean.(*User))
	//	if idx == 4 {
	//		usr = bean.(*User)
	//	}
	//	return nil
	//})
	//println(usr.Name)
	//if err != nil {
	//	println("Err")
	//	return
	//}

	////rows举例使用
	////Rows 和 Iterate 都是迭代数据库记录
	////Rows：相比Iterate更底层，更灵活
	////Iterate：高级的Rows封装
	//rows, err := engine.Where("id > ?", 2).Rows(new(User))
	//defer rows.Close()
	//if err != nil{
	//	fmt.Println("err:", err)
	//}
	////bean := &User{}
	//bean := new(User)
	////rows.Next()返回一个bool,是否有下一个数据
	//for rows.Next() {
	//	//Scan扫描记录
	//	if err := rows.Scan(bean);err!=nil {
	//		fmt.Println("err:", err)
	//	}else {
	//		fmt.Printf("%#v\n", bean)
	//	}
	//}

	//5.更新一条或者多条记录
	//var user User
	//user.Name = "ThreeUpdated"
	//affcted, err := engine.ID(8).Update(&user)
	//if err != nil {
	//	fmt.Printf("err:", err)
	//}
	//println(affcted)

	//6.删除一条或者多条记录，必须存在删除条件
	//affected, err := engine.ID(8).Delete(new(User))
	//if err !=nil {
	//	fmt.Printf("err:", err)
	//}
	//println(affected)

	//7.查询记录条数
	//counts, err := engine.Count(new(User))
	//if err != nil {
	//	fmt.Printf("err:", err)
	//}
	//println(counts)
}

//条件
func Query2() {
	engine := GetEngine()
	engine.ShowSQL(true)

	//1.Id, in:
	//var user User
	//has, err := engine.ID(1).Get(user)
	//has, err := engine.ID(core.PK{1, 2}).Get(user)	//用于复合主键
	//if err != nil {
	//	println(err)
	//	return
	//}
	//if has {
	//	fmt.Print(user)
	//}

	//var users []User
	////_ = engine.In("id", 1, 2, 3).Find(&users)
	//_ = engine.In("id", []int{1, 2, 3}).Find(&users)
	//for _, user := range users {
	//	fmt.Print(user)
	//}

	////2.Where, And, Or
	//var users []User
	//engine.Where("name = ?", "testInsert3").And("age = ?", 20).Or("age > ?", 20).Find(&users)
	//for _, user := range users {
	//	fmt.Print(user)
	//}
	//// SELECT * FROM user WHERE (.. AND ..) OR ...

	//3.OrderBy, Asc, Desc
	//var users []User
	////engine.Asc("age").Find(&users)
	//engine.Desc("age").Find(&users)
	//for _, user := range users {
	//	fmt.Println(user)
	//}

	////4.Limit, Top
	//var users = make([]User, 0)
	////engine.Limit(3, 0).Find(&users)
	//engine.
	//for _, user := range users {
	//	fmt.Println(user)
	//}

	////5.Sql,查询原生SQL
	//var users []User
	//engine.SQL("select * from prefix_user").Find(&users)
	//for _, user := range users {
	//	fmt.Println(user)
	//}

	////6.Cols, Omit, Distinct
	//var users []User
	//engine.Cols("name").Find(&users)
	////engine.Cols("name").Where("name = ?", "one").Update(&User{Name:"oneUpdated"})
	////engine.Omit("age").Find(&users)
	////engine.Omit("age").Insert(&User{Name:"insertTTT", Age: 66})
	////engine.Distinct("age").Find(&users)
	//for _, user := range users {
	//	fmt.Println(user)
	//}

	//7.Join, GroupBy, Having
	//var users []User
	//engine.GroupBy("name").Having("name='oneUpdated'").Find(&users)
	//for _, user := range users {
	//	fmt.Println(user)
	//}

	//一下没测试，没有两个表
	//var users = make([]User, 0)
	//engine.Join("LEFT", "userDetail", "user.id = userDetail.id").Find(&users)
	////SELECT * FROM user LEFT JOIN userdetail ON user.id=userdetail.id
}



//https://blog.csdn.net/feiwutudou/article/details/81317558
//https://blog.csdn.net/qq_769932247/article/details/84027445
//https://blog.csdn.net/qfzhangxu/article/details/89021741
//https://blog.csdn.net/qfzhangxu
//http://gobook.io/read/gitea.com/xorm/manual-en-US
