package model

import (
	"database/sql"
	"time"
)

// RateType RateType Model
type RateType struct {
	ID        int
	Name      string
	Rate      int
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
