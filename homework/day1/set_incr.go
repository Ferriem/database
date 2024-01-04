package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var Rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func main() {
	ctx := context.Background()
	err := Rdb.Set(ctx, "ferriem", "10", 0).Err()
	if err != nil {
		fmt.Println(err)
		return
	}
	val, err := Rdb.Get(ctx, "ferriem").Result()
	if err == nil {
		fmt.Println(val)
	} else {
		fmt.Println(err)
	}
	val1, err1 := Rdb.Incr(ctx, "ferriem").Result()
	if err1 == nil {
		fmt.Println(val1)
	} else {
		fmt.Println(err1)
	}
	val2, err2 := Rdb.IncrBy(ctx, "ferriem", 10).Result()
	if err2 == nil {
		fmt.Println(val2)
	} else {
		fmt.Println(err2)
	}
}
