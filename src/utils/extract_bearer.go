package utils

import (
	"strings"
)

// ExtractJWTFromBearerToken extracts the JWT string from a Bearer token.
func ExtractJWTFromBearerToken(bearerToken string) (string, bool) {
	// Check if the token starts with "Bearer "
	if strings.HasPrefix(bearerToken, "Bearer ") {
		// Split the token to get the JWT part
		parts := strings.Split(bearerToken, " ")
		if len(parts) == 2 {
			return parts[1], false // Return the JWT string and true if successfully extracted
		}
	}

	// If the format is not "Bearer <JWT>", return an empty string and false
	return "", true
}
