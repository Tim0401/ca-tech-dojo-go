package model

import "time"

// User ユーザーモデル
type User struct {
	ID        int16
	Name      string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
