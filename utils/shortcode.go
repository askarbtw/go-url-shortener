package utils

import (
	"crypto/rand"
	"math/big"
	"net/url"
	"strings"
)

const (
	// Characters used for generating short codes
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	// Length of the generated short code
	shortCodeLength = 6
)

// GenerateShortCode generates a random short code
func GenerateShortCode() (string, error) {
	shortCode := make([]byte, shortCodeLength)
	charsetLength := big.NewInt(int64(len(charset)))

	for i := 0; i < shortCodeLength; i++ {
		randomIndex, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", err
		}
		shortCode[i] = charset[randomIndex.Int64()]
	}

	return string(shortCode), nil
}

// ValidateURL performs validation on a URL
func ValidateURL(urlStr string) bool {
	// Basic length check
	if len(urlStr) == 0 {
		return false
	}

	// Ensure URL has a scheme (protocol)
	if !strings.Contains(urlStr, "://") {
		urlStr = "https://" + urlStr
	}

	// Parse URL
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return false
	}

	// Verify scheme
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return false
	}

	// Verify host
	if parsedURL.Host == "" {
		return false
	}

	return true
}

// PrepareURL ensures a URL has a protocol prefix and returns the prepared URL
func PrepareURL(urlStr string) string {
	// Ensure URL has a scheme (protocol)
	if !strings.Contains(urlStr, "://") {
		return "https://" + urlStr
	}
	return urlStr
}
