package internal

import (
	"encoding/json"
	"time"
)

type Job struct {
	ID          int64
	Kind        string
	Args        json.RawMessage
	Attempt     int
	MaxAttempts int
	Errors      []AttemptError
}

type AttemptError struct {
	Attempt int       `json:"attempt"`
	At      time.Time `json:"at"`
	Error   string    `json:"error"`
}
