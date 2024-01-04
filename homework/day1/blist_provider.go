package main

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var Rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func main() {
	ctx := context.Background()
	err := Rdb.RPush(ctx, "list1", "happy new year !").Err()
	if err != nil {
		panic(err)
	}
}
