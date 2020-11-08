package model

import (
	"database/sql"
	"time"
)

// Gacha Gacha Model
type Gacha struct {
	ID          int
	CharaID     int
	GachaTypeID int
	RateTypeID  int
	Rate        int
	CreatedAt   time.Time
	UpdatedAt   sql.NullTime
}
