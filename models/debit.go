package models

import "time"

type Debit struct {
	Title       string
	Description string
	Value       float64
	CreatedAt   time.Time
	PaymentID   string
	TypeID      int
	DateId      string
}
