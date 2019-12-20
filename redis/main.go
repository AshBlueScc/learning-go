package main

import (
	"time"
	"xorm/redis/redigo"
)

func main(){
	redigo.Insert("test-key", "test-value")
	redigo.Get("test-key")
	redigo.Exist("test-key")
	redigo.Get("test-key")
	redigo.DelKey("test-key")
	redigo.Exist("test-key")
	redigo.SetEx("test-key1", "test-value1", "5")

	time.Sleep(6 * time.Second)
	redigo.Exist("test-key1")

	lmap := map[string]string{"username":"666", "password":"666"}
	redigo.Json2Redis(lmap)

}