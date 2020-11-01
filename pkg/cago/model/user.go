package model

import (
	"database/sql"
	"time"
)

// User ユーザーモデル
type User struct {
	ID        int16
	Name      string
	Token     string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
