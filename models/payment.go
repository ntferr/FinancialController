package models

import "time"

type Payment struct {
	ID     string
	Type   string    `json:"type"`
	PaidAt time.Time `json:"paid_at"`
}
