package main

import "xorm/redis/go-redis"

//redigo test
//func main(){
//	redigo.Insert("test-key", "test-value")
//	redigo.Get("test-key")
//	redigo.Exist("test-key")
//	redigo.Get("test-key")
//	redigo.DelKey("test-key")
//	redigo.Exist("test-key")
//	redigo.SetEx("test-key1", "test-value1", "1")
//
//	time.Sleep(2 * time.Second)
//	redigo.Exist("test-key1")
//
//	lmap := map[string]string{"username":"666", "password":"666"}
//	redigo.Json2Redis(lmap)
//
//	redigo.Lpush("runoobkey","mysql", "mongodb", "redis")
//	redigo.PipDo()
//}


// go-redis test
func main(){
	redis.Pingr();
	redis.Setr("fool", "bar")
	redis.Getr("fool")
	redis.GetTtl("fool")

	redis.ListRpush("database", "mysql", "sqlite", "mongodb")
	redis.Hset("tmap", "key11", "value11")

	kv_map := make(map[string]interface{})
	kv_map["s3"] = "f3"
	kv_map["s4"] = "f4"
	redis.Hmset("tmap", kv_map)

	redis.Getmap("tmap", "s3")
	redis.GetAllMap("tmap")
	redis.Pubsub()
}