package config

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisCache represents a Redis connection
type RedisCache struct {
	Client *redis.Client
}

// ConnectRedis establishes a connection to Redis
func ConnectRedis(config *Config) *RedisCache {
	// Create a Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisURI,
		Password: config.RedisPassword,
		DB:       0, // default DB
	})

	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Printf("Warning: Failed to connect to Redis: %v", err)
		return nil
	}

	log.Println("Connected to Redis")
	return &RedisCache{
		Client: client,
	}
}

// Close disconnects from Redis
func (r *RedisCache) Close() {
	if r.Client != nil {
		if err := r.Client.Close(); err != nil {
			log.Printf("Error closing Redis connection: %v", err)
		} else {
			log.Println("Disconnected from Redis")
		}
	}
}
