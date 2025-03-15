package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/askarbtw/url-shortener-golang/models"
	"github.com/askarbtw/url-shortener-golang/services"
	"github.com/gorilla/mux"
)

// URLController handles HTTP requests for URL operations
type URLController struct {
	service *services.URLService
	baseURL string
}

// NewURLController creates a new instance of URLController
func NewURLController(service *services.URLService, baseURL string) *URLController {
	return &URLController{
		service: service,
		baseURL: baseURL,
	}
}

// CreateURL handles the creation of a new short URL
func (c *URLController) CreateURL(w http.ResponseWriter, r *http.Request) {
	var req models.CreateURLRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Create URL
	url, err := c.service.CreateURL(req.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create response
	response := models.URLResponse{
		ID:        url.ID,
		URL:       url.OriginalURL,
		ShortCode: url.ShortCode,
		CreatedAt: url.CreatedAt,
		UpdatedAt: url.UpdatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// GetURL retrieves a URL by its short code
func (c *URLController) GetURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortCode := vars["shortCode"]

	// Get URL
	url, err := c.service.GetURL(shortCode)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	// Create response
	response := models.URLResponse{
		ID:        url.ID,
		URL:       url.OriginalURL,
		ShortCode: url.ShortCode,
		CreatedAt: url.CreatedAt,
		UpdatedAt: url.UpdatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// RedirectURL redirects to the original URL
func (c *URLController) RedirectURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortCode := vars["shortCode"]

	// Get URL
	url, err := c.service.GetURL(shortCode)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	// Increment access count
	go c.service.IncrementAccessCount(shortCode)

	// Prepare the target URL
	targetURL := url.OriginalURL

	// Make sure we have a protocol prefix
	if !strings.HasPrefix(targetURL, "http://") && !strings.HasPrefix(targetURL, "https://") {
		targetURL = "https://" + targetURL
	}

	log.Printf("Redirecting %s to %s", shortCode, targetURL)

	// Redirect to the original URL
	http.Redirect(w, r, targetURL, http.StatusTemporaryRedirect)
}

// UpdateURL updates an existing URL
func (c *URLController) UpdateURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortCode := vars["shortCode"]

	var req models.UpdateURLRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Update URL
	url, err := c.service.UpdateURL(shortCode, req.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create response
	response := models.URLResponse{
		ID:        url.ID,
		URL:       url.OriginalURL,
		ShortCode: url.ShortCode,
		CreatedAt: url.CreatedAt,
		UpdatedAt: url.UpdatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DeleteURL deletes a URL
func (c *URLController) DeleteURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortCode := vars["shortCode"]

	// Delete URL
	err := c.service.DeleteURL(shortCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetURLStats retrieves statistics for a URL
func (c *URLController) GetURLStats(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortCode := vars["shortCode"]

	// Get URL
	url, err := c.service.GetURL(shortCode)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	// Create response
	response := models.URLStatsResponse{
		ID:          url.ID,
		URL:         url.OriginalURL,
		ShortCode:   url.ShortCode,
		CreatedAt:   url.CreatedAt,
		UpdatedAt:   url.UpdatedAt,
		AccessCount: url.AccessCount,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetAllURLStats retrieves statistics for all URLs
func (c *URLController) GetAllURLStats(w http.ResponseWriter, r *http.Request) {
	// Get all URLs with stats
	urls, err := c.service.GetAllURLsWithStats()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create response
	var response []models.URLStatsResponse
	for _, url := range urls {
		response = append(response, models.URLStatsResponse{
			ID:          url.ID,
			URL:         url.OriginalURL,
			ShortCode:   url.ShortCode,
			CreatedAt:   url.CreatedAt,
			UpdatedAt:   url.UpdatedAt,
			AccessCount: url.AccessCount,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
