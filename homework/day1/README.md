```sh
go run set_incr.go  
go run blist_digester.go
go run blist_provider.go
```

```go
//connect to redis
var Rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

ctx := context.Background()
```

