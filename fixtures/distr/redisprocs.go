package main

import (
	"github.com/go-redis/redis"
	"log"
)

type RedisDistr struct {
	client *redis.Client
}

func NewRedisDistr() RedisDistr {
	return RedisDistr{client: redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})}
}

func (c *RedisDistr) Ping() {
	log.Println(".. ping redis")

	pong, err := c.client.Ping().Result()
	log.Println(pong, err)
}
