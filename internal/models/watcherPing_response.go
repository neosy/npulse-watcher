package models

type WatcherPingResponse struct {
	Success bool   `json:"success"`
	Text    string `json:"text"`
}
