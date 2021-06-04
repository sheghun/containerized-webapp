package memory

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var ErrKeyDoesNotExist = errors.New("key does not exist")

func redisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "",
		DB:       0,
	})

}

// Get retrieves the key from memory
func Get(ctx context.Context, key string) (string, error) {
	fmt.Println(os.Getenv("REDIS_URL"))
	val, err := redisClient().Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", ErrKeyDoesNotExist

		}
		log.Printf("error occured trying to retreive key from redis: %v", err)
		return "", err
	}

	return val, nil
}

// Set saves the key to the memory
func Set(ctx context.Context, key string, value string) error {
	return redisClient().Set(ctx, key, value, 1*time.Hour).Err()
}
