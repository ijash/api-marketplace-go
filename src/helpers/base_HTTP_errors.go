package helpers

import (
	"ijash-jwt-auth/src/models"
)

// HTTPErrorUnauthorized creates a models.BaseHTTPError for Unauthorized (401).
func Unauthorized(message ...string) models.BaseHTTPError {
	msg := "Unauthorized" // Default message if not provided
	if len(message) > 0 {
		msg = message[0] // Use the provided message if available
	}

	return models.BaseHTTPError{
		Status:  false,
		Message: msg,
	}
}

// HTTPErrorInternalServerError creates a models.BaseHTTPError for Internal Server Error (500).
func InternalServerError(message ...string) models.BaseHTTPError {
	msg := "Internal Server Error" // Default message if not provided
	if len(message) > 0 {
		msg = message[0] // Use the provided message if available
	}

	return models.BaseHTTPError{
		Status:  false,
		Message: msg,
	}
}

// HTTPErrorBadRequest creates a models.BaseHTTPError for Bad Request (400).
func BadRequest(message ...string) models.BaseHTTPError {
	msg := "Bad Request" // Default message if not provided
	if len(message) > 0 {
		msg = message[0] // Use the provided message if available
	}

	return models.BaseHTTPError{
		Status:  false,
		Message: msg,
	}
}

// HTTPErrorNotFound creates a models.BaseHTTPError for Not Found (404).
func NotFound(message ...string) models.BaseHTTPError {
	msg := "Not Found" // Default message if not provided
	if len(message) > 0 {
		msg = message[0] // Use the provided message if available
	}

	return models.BaseHTTPError{
		Status:  false,
		Message: msg,
	}
}
