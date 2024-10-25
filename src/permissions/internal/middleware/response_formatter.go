package middleware

import "time"

type MessageFormat struct {
	Message   string `json:"message"`
	Data      any    `json:"data"`
	Timestamp string `json:"timestamp"`
}

func Format(message string, data any) MessageFormat {
	currentTime := time.Now().UTC().Format(time.RFC3339)

	return MessageFormat{
		Message:   message,
		Data:      data,
		Timestamp: currentTime,
	}
}
