package main

import (
	"github.com/go-redis/redis"
	"log"
)

func main() {
	rdclient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := rdclient.Ping().Result()
	if err != nil {
		panic(err)
	}
	log.Println("redis connect success")
	task := "111"
	
}
