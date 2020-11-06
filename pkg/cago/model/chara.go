package model

import (
	"database/sql"
	"time"
)

// Gacha Gacha Model
type Chara struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
