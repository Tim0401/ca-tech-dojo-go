package model

import (
	"database/sql"
	"time"
)

// Chara Chara Model
type Chara struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type CharaUser struct {
	ID        int
	UserID    int
	CharaID   int
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
