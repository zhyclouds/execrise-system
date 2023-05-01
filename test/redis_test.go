package test

import (
	"context"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func TestRedisGet(t *testing.T) {
	rdb.Set(ctx, "key", "value", time.Second*20)
}

func TestRedisSet(t *testing.T) {
	result, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
