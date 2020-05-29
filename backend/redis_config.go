package backend

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
)

func RClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ping(client)

	return client
}

func ping(client *redis.Client) error {

	ctx := context.Background()
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		return err
	}
	fmt.Println(pong, err)
	// Output: PONG <nil>

	return nil
}