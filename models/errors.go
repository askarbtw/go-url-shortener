package models

import "errors"

// Application errors
var (
	ErrorInvalidURL          = errors.New("invalid URL")
	ErrorGeneratingShortCode = errors.New("failed to generate unique short code")
	ErrorURLNotFound         = errors.New("URL not found")
	ErrorShortCodeExists     = errors.New("short code already exists")
)
