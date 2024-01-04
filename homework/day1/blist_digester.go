package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var Rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func main() {
	ctx := context.Background()
	val, err := Rdb.BRPop(ctx, 300*time.Second, "list1").Result()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(val)
	}
}
