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
	zMembers1 := []redis.Z{
		{Score: 500, Member: "7wks"},
		{Score: 9, Member: "gog"},
		{Score: 9999, Member: "prag"},
	}
	err := Rdb.Del(ctx, "visits").Err()
	if err != nil {
		panic(err)
	}
	err = Rdb.Del(ctx, "votes").Err()
	if err != nil {
		panic(err)
	}
	err = Rdb.Del(ctx, "importance").Err()
	if err != nil {
		panic(err)
	}
	err = Rdb.ZAdd(ctx, "visits", zMembers1...).Err()
	if err != nil {
		panic(err)
	}
	zMembers2 := []redis.Z{
		{Score: 2, Member: "7wks"},
		{Score: 0, Member: "gog"},
		{Score: 9001, Member: "prag"},
	}
	err = Rdb.ZAdd(ctx, "votes", zMembers2...).Err()
	if err != nil {
		panic(err)
	}
	err = Rdb.ZUnionStore(ctx, "importance", &redis.ZStore{Keys: []string{"visits", "votes"}, Weights: []float64{1, 2}, Aggregate: "SUM"}).Err()
	if err != nil {
		panic(err)
	}
	zrange := Rdb.ZRangeWithScores(ctx, "importance", 0, -1).Val()
	var zSetMembers []redis.Z
	for _, z := range zrange {
		zSetMembers = append(zSetMembers, redis.Z{Member: z.Member, Score: z.Score})
	}
	fmt.Println("Sorted Set Members:")
	for _, member := range zSetMembers {
		fmt.Printf("Member: %s, Score: %v\n", member.Member, member.Score)
	}
	Rdb.Close()

}
