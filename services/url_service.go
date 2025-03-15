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
}

// NewURLService creates a new instance of URLService
func NewURLService(repository *repositories.URLRepository) *URLService {
	return &URLService{
		repository: repository,
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
			return createdURL, nil
		}
		log.Printf("Failed to create URL with short code %s (attempt %d): %v", shortCode, attempt+1, err)
	}

	// If we couldn't create a unique short code after multiple attempts
	return models.URL{}, models.ErrorGeneratingShortCode
}

// GetURL retrieves a URL by its short code
func (s *URLService) GetURL(shortCode string) (models.URL, error) {
	return s.repository.GetURLByShortCode(shortCode)
}

// UpdateURL updates an existing URL
func (s *URLService) UpdateURL(shortCode string, originalURL string) (models.URL, error) {
	// Validate URL
	if !utils.ValidateURL(originalURL) {
		return models.URL{}, models.ErrorInvalidURL
	}

	// Ensure URL has proper protocol prefix
	originalURL = utils.PrepareURL(originalURL)

	return s.repository.UpdateURL(shortCode, originalURL)
}

// DeleteURL deletes a URL
func (s *URLService) DeleteURL(shortCode string) error {
	return s.repository.DeleteURL(shortCode)
}

// IncrementAccessCount increments the access count for a URL
func (s *URLService) IncrementAccessCount(shortCode string) error {
	return s.repository.IncrementAccessCount(shortCode)
}
