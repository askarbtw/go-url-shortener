package repositories

import (
	"context"
	"log"
	"time"

	"github.com/askarbtw/url-shortener-golang/config"
	"github.com/askarbtw/url-shortener-golang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// URLRepository handles database operations for URLs
type URLRepository struct {
	collection *mongo.Collection
}

// NewURLRepository creates a new instance of URLRepository
func NewURLRepository(db *config.Database) *URLRepository {
	// Create a unique index on short_code to prevent duplicates
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "short_code", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := db.DB.Collection("urls").Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Printf("Warning: Failed to create unique index on short_code: %v", err)
	}

	return &URLRepository{
		collection: db.DB.Collection("urls"),
	}
}

// CreateURL creates a new URL in the database
func (r *URLRepository) CreateURL(url models.URL) (models.URL, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if short code already exists
	var existingURL models.URL
	err := r.collection.FindOne(ctx, bson.M{"short_code": url.ShortCode}).Decode(&existingURL)
	if err == nil {
		return models.URL{}, models.ErrorShortCodeExists
	} else if err != mongo.ErrNoDocuments {
		log.Printf("Error checking for existing short code: %v", err)
		return models.URL{}, err
	}

	url.CreatedAt = time.Now()
	url.UpdatedAt = time.Now()
	url.AccessCount = 0

	result, err := r.collection.InsertOne(ctx, url)
	if err != nil {
		// Check if the error is due to a duplicate key (race condition)
		if mongo.IsDuplicateKeyError(err) {
			return models.URL{}, models.ErrorShortCodeExists
		}
		log.Printf("Error inserting URL: %v", err)
		return models.URL{}, err
	}

	url.ID = result.InsertedID.(primitive.ObjectID)
	return url, nil
}

// GetURLByShortCode retrieves a URL by its short code
func (r *URLRepository) GetURLByShortCode(shortCode string) (models.URL, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var url models.URL
	err := r.collection.FindOne(ctx, bson.M{"short_code": shortCode}).Decode(&url)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.URL{}, models.ErrorURLNotFound
		}
		return models.URL{}, err
	}

	return url, nil
}

// UpdateURL updates a URL in the database
func (r *URLRepository) UpdateURL(shortCode string, originalURL string) (models.URL, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var url models.URL
	err := r.collection.FindOne(ctx, bson.M{"short_code": shortCode}).Decode(&url)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.URL{}, models.ErrorURLNotFound
		}
		return models.URL{}, err
	}

	update := bson.M{
		"$set": bson.M{
			"original_url": originalURL,
			"updated_at":   time.Now(),
		},
	}

	_, err = r.collection.UpdateOne(ctx, bson.M{"short_code": shortCode}, update)
	if err != nil {
		return models.URL{}, err
	}

	// Fetch updated document
	err = r.collection.FindOne(ctx, bson.M{"short_code": shortCode}).Decode(&url)
	if err != nil {
		return models.URL{}, err
	}

	return url, nil
}

// DeleteURL deletes a URL from the database
func (r *URLRepository) DeleteURL(shortCode string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := r.collection.DeleteOne(ctx, bson.M{"short_code": shortCode})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return models.ErrorURLNotFound
	}

	return nil
}

// IncrementAccessCount increments the access count for a URL
func (r *URLRepository) IncrementAccessCount(shortCode string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{
		"$inc": bson.M{"access_count": 1},
	}

	result, err := r.collection.UpdateOne(ctx, bson.M{"short_code": shortCode}, update)
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return models.ErrorURLNotFound
	}

	return nil
}
