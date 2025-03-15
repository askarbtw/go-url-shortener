package services

import (
	"log"

	"github.com/askarbtw/url-shortener-golang/models"
	"github.com/askarbtw/url-shortener-golang/repositories"
	"github.com/askarbtw/url-shortener-golang/utils"
)

// URLService handles business logic for URL operations
type URLService struct {
	repository *repositories.URLRepository
	cache      *CacheService
}

// NewURLService creates a new instance of URLService
func NewURLService(repository *repositories.URLRepository, cache *CacheService) *URLService {
	return &URLService{
		repository: repository,
		cache:      cache,
	}
}

// CreateURL creates a new short URL
func (s *URLService) CreateURL(originalURL string) (models.URL, error) {
	// Validate URL
	if !utils.ValidateURL(originalURL) {
		return models.URL{}, models.ErrorInvalidURL
	}

	// Ensure URL has proper protocol prefix
	originalURL = utils.PrepareURL(originalURL)

	// Generate a unique short code
	var shortCode string
	var err error
	var url models.URL
	maxAttempts := 10

	// Try multiple times to generate a unique short code
	for attempt := 0; attempt < maxAttempts; attempt++ {
		shortCode, err = utils.GenerateShortCode()
		if err != nil {
			log.Printf("Error generating short code (attempt %d): %v", attempt+1, err)
			continue
		}

		// Create URL object
		url = models.URL{
			OriginalURL: originalURL,
			ShortCode:   shortCode,
		}

		// Try to save to database
		createdURL, err := s.repository.CreateURL(url)
		if err == nil {
			// Store in cache
			if s.cache != nil {
				s.cache.SetURL(createdURL)
			}
			return createdURL, nil
		}
		log.Printf("Failed to create URL with short code %s (attempt %d): %v", shortCode, attempt+1, err)
	}

	// If we couldn't create a unique short code after multiple attempts
	return models.URL{}, models.ErrorGeneratingShortCode
}

// GetURL retrieves a URL by its short code
func (s *URLService) GetURL(shortCode string) (models.URL, error) {
	// Try to get from cache first
	if s.cache != nil {
		if url, found := s.cache.GetURL(shortCode); found {
			log.Printf("Cache hit for shortCode: %s", shortCode)
			return url, nil
		}
	}

	// If not in cache, get from database
	url, err := s.repository.GetURLByShortCode(shortCode)
	if err != nil {
		return models.URL{}, err
	}

	// Store in cache for future requests
	if s.cache != nil {
		s.cache.SetURL(url)
	}

	return url, nil
}

// UpdateURL updates an existing URL
func (s *URLService) UpdateURL(shortCode string, originalURL string) (models.URL, error) {
	// Validate URL
	if !utils.ValidateURL(originalURL) {
		return models.URL{}, models.ErrorInvalidURL
	}

	// Ensure URL has proper protocol prefix
	originalURL = utils.PrepareURL(originalURL)

	// Update in database
	updatedURL, err := s.repository.UpdateURL(shortCode, originalURL)
	if err != nil {
		return models.URL{}, err
	}

	// Update cache
	if s.cache != nil {
		s.cache.SetURL(updatedURL)
	}

	return updatedURL, nil
}

// DeleteURL deletes a URL
func (s *URLService) DeleteURL(shortCode string) error {
	// Delete from database
	err := s.repository.DeleteURL(shortCode)
	if err != nil {
		return err
	}

	// Invalidate cache
	if s.cache != nil {
		s.cache.InvalidateURL(shortCode)
	}

	return nil
}

// IncrementAccessCount increments the access count for a URL
func (s *URLService) IncrementAccessCount(shortCode string) error {
	// Increment in database
	err := s.repository.IncrementAccessCount(shortCode)
	if err != nil {
		return err
	}

	// Invalidate cache since the access count has changed
	if s.cache != nil {
		s.cache.InvalidateURL(shortCode)
	}

	return nil
}

// GetAllURLsWithStats retrieves all URLs with their statistics
func (s *URLService) GetAllURLsWithStats() ([]models.URL, error) {
	return s.repository.GetAllURLs()
}
