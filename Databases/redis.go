package Databases

import (
	"fmt"
	"github.com/go-redis/redis"
)

var RedisDB *redis.Client

func init(){
	RedisDB = redis.NewClient(&redis.Options{
		Addr:string("127.0.0.1:6379"),
		DB:0,
		Password:"",
	})

	_,err := RedisDB.Ping().Result()
	if err != nil{
		panic(err)
	}
	fmt.Println("Redis-连接成功")
}