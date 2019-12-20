package redigo

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//插入值
func Insert(key string, value string){
	_, err := Conn.Do("SET", key, value)
	if err != nil {
		fmt.Println("redis set value failed >>>", err)
		return
	}
	fmt.Printf("Operation: SET %s %s \n",key,value)
}

//检验key值是否存在
func Exist(key string) bool{
	exists, err := redis.Bool(Conn.Do("EXISTS", key))
	if err != nil {
		fmt.Printf("illegal exception %s \n", err)
	}
	fmt.Printf("exists or not: %v \n", exists)
	return exists
}

//获取值
func Get(key string) string {
	value, err := redis.String(Conn.Do("GET", key))
	//value, err := conn.Do("GET", "test-key")		//这个返回的是一个interface
	if err != nil {
		fmt.Println("redis get value failed >>>", err)
		return ""
	}
	fmt.Printf("Operation:Get %s: %s \n", key, value)
	return value
}

//给定一个kv的过期时间
//EX secondes, 5秒
//PX milliseconds
func SetEx(key string, value string, seconds string){
	_, err := Conn.Do("SET", key, value, "EX", seconds)
	if err != nil {
		fmt.Println("redis set value failed >>>", err)
	}
}

//删除key
func DelKey(key string){
	_, err := Conn.Do("DEL", key)
	if err != nil {
		fmt.Println("redis delete value failed >>>", err)
	}
}

//读写json到redis
func Json2Redis(jsonString map[string]string){
	key := "profile"
	value, _ := json.Marshal(jsonString)

	judge, err := Conn.Do("SETNX", key, value)
	if err != nil {
		fmt.Println(err)
	}
	if judge == int64(1) {
		fmt.Println("success")
	}

	var imapGet map[string]string

	valueGet, err := redis.Bytes(Conn.Do("Get", key))
	if err != nil {
		fmt.Println(err)
	}

	errShal := json.Unmarshal(valueGet, &imapGet)
	if errShal != nil {
		fmt.Println(err)
	}

	fmt.Println(imapGet["username"])
	fmt.Println(imapGet["password"])
}

//lpush
func Lpush(key string, values ...string)(){
	for _, value := range values{
		_, err := Conn.Do("lpush", key, value)
		if err != nil {
			fmt.Printf("redis set failed: %s", err)
			return
		}
	}

	vs, _ := redis.Values(Conn.Do("lrange", key, "0", "100"))

	for _, v := range vs{
		fmt.Println(string(v.([]byte)))
	}

}

//管道
func PipDo()(){
	Conn.Send("SET", "foo", "bar")
	Conn.Send("Get", "foo")
	Conn.Flush()
	Conn.Receive()
	v, _ := Conn.Receive()
	fmt.Println(string(v.([]byte)))
}