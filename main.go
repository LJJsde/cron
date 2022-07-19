package main

import (
	"awesomeProject7/service"
	"github.com/go-redis/redis"
	"log"
	"time"
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
	service.Create(1).Weekday(1, 0).Loc(time.Local).Run(task)
}
