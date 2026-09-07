package model

import (
	"context"
	"log"
	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client
var Ctx = context.Background()

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	if err := RDB.Ping(Ctx).Err(); err != nil{
		log.Fatalf("Redis connection failed, %v", err)
	}
	log.Println("Redis connection successed")
}