package model

import (
	"database/sql"
	"time"
)

// User ユーザーモデル
type User struct {
	ID        int
	Name      string
	Token     string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type key int

const (
	UserKey key = iota
)
