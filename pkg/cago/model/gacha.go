package model

import (
	"database/sql"
	"time"
)

// GachaProbability GachaProbability Model
type GachaProbability struct {
	GroupID   string
	Number    int
	CharaID   int
	Rate      int
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type GachaProbabilityGroup struct {
	GachaTypeID             int
	Number                  int
	GachaProbabilityGroupID string
	Rate                    int
	CreatedAt               time.Time
	UpdatedAt               sql.NullTime
}
