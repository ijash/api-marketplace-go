package helpers

import (
	"ijash-jwt-auth/src/models"
)

func BaseResponseOk(data interface{}, message ...string) models.BaseResponse {
	msg := "OK" // Default message if not provided
	if len(message) > 0 {
		msg = message[0] // Use the provided message if available
	}

	return models.BaseResponse{
		Status:  true,
		Message: msg,
		Data:    data,
	}
}
