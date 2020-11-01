package model

import (
	"database/sql"
	"time"
)

// Gacha Gacha Model
type Gacha struct {
	ID        int32
	CharaID   int32
	Rate      float64
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
