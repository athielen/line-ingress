package backend

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"line-ingress/models"
)

type RedisBackend struct {
	Client *redis.Client
	Lines  chan models.Line
}

func (rb *RedisBackend) Consume() {
	fmt.Println("Register the worker")
	for i := range rb.Lines {
		fmt.Println("redis..." + i.String() + "\n")

		ctx := context.Background()

		rb.Client.Set(ctx,"key-02", i.String(), 0).Result()

		val2 := rb.Client.Get(ctx, "key-02")

		fmt.Println(val2.String())

	}
}

func (rb *RedisBackend) Input() chan<- models.Line {
	return rb.Lines
}

func (rb *RedisBackend) Output() <-chan models.Line {
	return rb.Lines
}
