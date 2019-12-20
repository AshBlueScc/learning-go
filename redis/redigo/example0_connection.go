package redigo

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var Conn redis.Conn

//Init函数会自调用
func init(){
	var err error
	if Conn == nil {
		Conn, err = redis.Dial("tcp",
			"127.0.0.1:6379",
			redis.DialDatabase(1), //dialOption参数可以配置选择数据库，连接密码，心跳检测等等
			redis.DialPassword(""))
		if err != nil {
			fmt.Println("connect to redis failed, caused by >>>", err)
			return
		}
	}
}

//为了方便改写一下
//创建连接
//func Conn() *redis.Conn{
//	conn, err := redis.Dial("tcp",
//				"127.0.0.1:6379",
//				redis.DialDatabase(1),//dialOption参数可以配置选择数据库，连接密码，心跳检测等等
//				redis.DialPassword(""))
//	if err != nil {
//		fmt.Println("connect to redis failed, caused by >>>", err)
//		return nil
//	}
//
//	return &conn
//}