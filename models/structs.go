package models

import "time"

type InfoLogs struct {
	Msg      string `json:"msg"`
	Metadata struct {
		SessionID  string `json:"session_id"`
		InternalID string `json:"internal_id"`
		MarketCode string `json:"market_code"`
		Images     []struct {
			Position int    `json:"position"`
			URL      string `json:"url"`
		} `json:"images"`
		ImageType   string `json:"image_type"`
		ProcessedID int    `json:"processed_id"`
	} `json:"metadata"`
	Version     string    `json:"@version"`
	Tags        []string  `json:"tags"`
	Application string    `json:"application"`
	Sender      string    `json:"sender"`
	ReceivedAt  time.Time `json:"received_at"`
	Level       string    `json:"level"`
	Timestamp   time.Time `json:"@timestamp"`
	Type        string    `json:"type"`
}


type Message struct {
	Msg       string `json:"msg"`
	TimeStamp time.Time `json:"@timestamp"`
	Application string `json:"application"`
}
