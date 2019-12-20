package redis

import (
	"fmt"
	"time"
)

//测试联通
func Pingr() () {
	pong, err := Client.Ping().Result()
	if err != nil {
		fmt.Printf("ping error[%s]\n", err.Error())
	}
	fmt.Printf("ping result: %s\n", pong)
}

//set test
func Setr(key string, value string) () {
	err := Client.Set(key, value, 0).Err()
	if err != nil {
		fmt.Printf("SET key[%s], failed err: %s\n", key, err)
	}
	fmt.Printf("Operation: SET %s %s \n", key, value)
}

//get test
func Getr(key string) (string) {
	value, err := Client.Get(key).Result()
	if err != nil {
		fmt.Printf("GET key[%s], failed err: %s\n", key, err)
	}
	fmt.Printf("Operation: GET %s: %s \n", key, value)
	return value
}

//get ttl(键到期时间)
func GetTtl(key string) () {
	duration, err := Client.TTL(key).Result()
	if err != nil {
		fmt.Printf("GET TTl key[%s]: err: %s\n", key, err)
	}
	fmt.Printf("key[%s]'s ttl is [%s] %ds\n", key, duration.String(), int64(duration.Seconds()))
}

//list test
func ListRpush(key string, values ...string) () {
	for _, value := range values {
		err := Client.RPush(key, value).Err()
		listLen, err1 := Client.LLen(key).Result()
		if err != nil {
			fmt.Printf("rpush list[%s], err: %s\n", key, err)
		}
		if err1 != nil {
			fmt.Printf("get len of list[%s], err: %s\n", key, err1)
		}
		fmt.Printf("list len: %d\n", listLen)
	}
	result, err3 := Client.BLPop(time.Second*1, key).Result()
	if err3 != nil {
		fmt.Printf("blpop list[%s] error: %s\n", key, err3)
	}
	fmt.Printf("blpop list[%s], value:%s \n", key, result[1])
}

//hmap
//第一个key表示这个set集合的key, 后面的key, value表示这个集合里面的键值对
func Hset(setKey string, key string, value string) () {
	err := Client.HSet(setKey, key, value).Err()
	if err != nil {
		fmt.Printf("hset map[%s] field[%s] value[%s] err:[%s]\n", setKey, key, value, err)
	}
}

//hmset
func Hmset(key string, kvs map[string]interface{}) () {
	err := Client.HMSet(key, kvs).Err()
	if err != nil {
		fmt.Printf("Hmset failed! Err: %s\n", err)
	}
	map_len, _ := Client.HLen(key).Result()
	fmt.Printf("len of map[%s]: %d\n", key, map_len)
}

//get map value
func Getmap(setKey string, key string)(){
	value, err := Client.HGet(setKey, key).Result()
	if err != nil {
		fmt.Printf("HGet failed!Error: %s\n", err)
	}
	fmt.Printf("field[%s] value of map[%s] is %s \n", key, setKey, value)
}

//hgetall
func GetAllMap(setKey string)(){
	result_kv, err := Client.HGetAll(setKey).Result()
	if err != nil {
		fmt.Printf("HGetAll failed!Error: %s\n", err)
	}

	for k, v := range result_kv {
		fmt.Printf("map[%s] %s = %s\n", setKey, k, v)
	}
}

//pubsub test
func Pubsub()(){
	pubsub := Client.Subscribe("test_channel")
	pubsub.Receive()

	ch := pubsub.Channel()
	Client.Publish("test_channel", "hello")

	time.AfterFunc(time.Second * 2, func() {
		pubsub.Close()
	})

	//consume message
	for {
		msg, ok := <-ch
		if !ok {
			break
		}

		fmt.Printf("recv message[%s] from channel[%s]\n", msg.Payload, msg.Channel)
	}
}