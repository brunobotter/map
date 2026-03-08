package domain

import "time"

type Event struct {
	Id        string    `json:"id"`
	Type      string    `json:"name"`
	Severity  int       `json:"severity"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Source    string    `json:"source"`
	CreatedAt time.Time `json:"created_at"`
}
