package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7000", ":7001", ":7002"},
	})
	ctx := context.Background()
	err := rdb.ForEachShard(ctx, func(ctx context.Context, shard *redis.Client) error {
		return shard.Ping(ctx).Err()
	})
	if err != nil {
		log.Fatal(err)
	}
	err = rdb.Set(ctx, "ferriem", "123", 0).Err()
	if err != nil {
		log.Fatal(err)
	}
	err = rdb.Set(ctx, "ferriem1", "123", 0).Err()
	if err != nil {
		log.Fatal(err)
	}
	err = rdb.Set(ctx, "ferriem2", "123", 0).Err()
	if err != nil {
		log.Fatal(err)
	}
	var value string
	value, err = rdb.Get(ctx, "ferriem").Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(value)
	rdb.Close()
}
