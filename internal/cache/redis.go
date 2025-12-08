// Package cache contains the redis setup for caching device resources
package cache

import (
	"context"
	"fmt"

	"synf/internal/api/data/models"

	"github.com/redis/go-redis/v9"
)

func Init() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // No password set
		DB:       0,  // Use default DB
		Protocol: 2,  // Connection protocol
	})

	return client
}

func Test(client *redis.Client) {
	ctx := context.Background()
	err := client.Set(ctx, "temperature", "44", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, "temperature").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(val)
}

func (cpu *models.Cpu) Store() {
}
