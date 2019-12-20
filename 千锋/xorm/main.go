package main

import (
	mysql2 "xorm/千锋/xorm/mysql"
)

func main() {
	//1.测试Connection，用的自己的一个表t_category
	//mysql.ConnectionTest()

	//2.测试Create,新建prefix_user和suffix_user
	//mysql.CreateTablesTest()

	//3.测试Insert，插入两个user到prefix_user表格
	//user := new(mysql.User)
	//user.Name = "testInsert3"
	//user.Age = 20
	//
	//user1 := new(mysql.User)
	//user1.Name = "testInsert3"
	//user1.Age = 21
	//mysql.Insert(user)
	//mysql.Insert(user1)

	//4.测试sql查询prefix表格
	//mysql.Query()

	//5.测试orm方法查询prefix_user表格
	mysql2.Query2()
}