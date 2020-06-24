package backend

import (
	"context"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"os"
)

func RClient() *redis.Client {
	log.Debug("Initializing redis client...")

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
		log.Fatal("Cannot connect to redis at " + client.Options().Addr + ". Gracefully shutting down...")
		os.Exit(1)
		return err
	}

	log.Debug("Successfully recieved PONG from redis server at" + client.Options().Addr + ":" + pong);

	return nil
}