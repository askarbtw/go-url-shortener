package services

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/askarbtw/url-shortener-golang/config"
	"github.com/askarbtw/url-shortener-golang/models"
)

// CacheService handles caching of URL data
type CacheService struct {
	cache    *config.RedisCache
	cacheTTL time.Duration
}

// NewCacheService creates a new instance of CacheService
func NewCacheService(cache *config.RedisCache, ttlSeconds int) *CacheService {
	return &CacheService{
		cache:    cache,
		cacheTTL: time.Duration(ttlSeconds) * time.Second,
	}
}

// GetURL tries to retrieve a URL from the cache
func (s *CacheService) GetURL(shortCode string) (models.URL, bool) {
	// If Redis is not connected, return not found
	if s.cache == nil || s.cache.Client == nil {
		return models.URL{}, false
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// Try to get the URL from cache
	key := "url:" + shortCode
	data, err := s.cache.Client.Get(ctx, key).Result()
	if err != nil {
		return models.URL{}, false
	}

	// Unmarshal the JSON data
	var url models.URL
	if err := json.Unmarshal([]byte(data), &url); err != nil {
		log.Printf("Error unmarshaling URL data from cache: %v", err)
		return models.URL{}, false
	}

	return url, true
}

// SetURL stores a URL in the cache
func (s *CacheService) SetURL(url models.URL) {
	// If Redis is not connected, do nothing
	if s.cache == nil || s.cache.Client == nil {
		return
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// Marshal the URL to JSON
	data, err := json.Marshal(url)
	if err != nil {
		log.Printf("Error marshaling URL data for cache: %v", err)
		return
	}

	// Set the URL in cache
	key := "url:" + url.ShortCode
	if err := s.cache.Client.Set(ctx, key, data, s.cacheTTL).Err(); err != nil {
		log.Printf("Error setting URL in cache: %v", err)
	}
}

// InvalidateURL removes a URL from the cache
func (s *CacheService) InvalidateURL(shortCode string) {
	// If Redis is not connected, do nothing
	if s.cache == nil || s.cache.Client == nil {
		return
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// Delete the URL from cache
	key := "url:" + shortCode
	if err := s.cache.Client.Del(ctx, key).Err(); err != nil {
		log.Printf("Error removing URL from cache: %v", err)
	}
}
