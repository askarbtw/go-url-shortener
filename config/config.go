package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
	MongoURI      string
	DBName        string
	Port          string
	BaseURL       string
	RedisURI      string
	RedisPassword string
	CacheTTL      int // Time to live for cached items in seconds
}

// LoadConfig loads the application configuration from environment variables
func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file:", err)
	}

	// Try to parse cache TTL, default to 3600 seconds (1 hour)
	cacheTTL := 3600
	if ttlStr := os.Getenv("CACHE_TTL"); ttlStr != "" {
		if ttl, err := strconv.Atoi(ttlStr); err == nil {
			cacheTTL = ttl
		} else {
			log.Printf("Warning: Invalid CACHE_TTL value '%s', using default: %d", ttlStr, cacheTTL)
		}
	}

	return &Config{
		MongoURI:      getEnv("MONGO_URI", "mongodb://localhost:27017"),
		DBName:        getEnv("DB_NAME", "url_shortener"),
		Port:          getEnv("PORT", "8080"),
		BaseURL:       getEnv("BASE_URL", "http://localhost:8080/"),
		RedisURI:      getEnv("REDIS_URI", "localhost:6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		CacheTTL:      cacheTTL,
	}
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
