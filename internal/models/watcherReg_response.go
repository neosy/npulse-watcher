package models

type WatcherRegSuccessResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}
