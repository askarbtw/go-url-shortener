package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// URL represents a URL shortening record
type URL struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	OriginalURL string             `json:"url" bson:"original_url"`
	ShortCode   string             `json:"shortCode" bson:"short_code"`
	AccessCount int                `json:"accessCount" bson:"access_count"`
	CreatedAt   time.Time          `json:"createdAt" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updatedAt" bson:"updated_at"`
}

// CreateURLRequest is used to parse the request for creating a URL
type CreateURLRequest struct {
	URL string `json:"url"`
}

// UpdateURLRequest is used to parse the request for updating a URL
type UpdateURLRequest struct {
	URL string `json:"url"`
}

// URLResponse represents the response object for a URL
type URLResponse struct {
	ID        primitive.ObjectID `json:"id"`
	URL       string             `json:"url"`
	ShortCode string             `json:"shortCode"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
}

// URLStatsResponse represents the response object for URL statistics
type URLStatsResponse struct {
	ID          primitive.ObjectID `json:"id"`
	URL         string             `json:"url"`
	ShortCode   string             `json:"shortCode"`
	CreatedAt   time.Time          `json:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt"`
	AccessCount int                `json:"accessCount"`
}
