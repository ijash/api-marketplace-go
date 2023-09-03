package models

type BaseHTTPError struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
