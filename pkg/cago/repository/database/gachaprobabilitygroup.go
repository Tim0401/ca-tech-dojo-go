package database

import (
	"ca-tech-dojo-go/pkg/cago/model"
	"database/sql"

	"golang.org/x/xerrors"
)

type dbGachaProbabilityGroupRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func (r *dbGachaProbabilityGroupRepository) FindByGachaType(gachaTypeID int) ([]model.GachaProbabilityGroup, error) {
	var gachaProbabilityGroups []model.GachaProbabilityGroup
	var rows *sql.Rows
	var err error
	cmd := "SELECT gacha_type_id, number, gacha_probability_group_id, rate, created_at, updated_at " +
		"FROM gacha_probability_group WHERE gacha_type_id = ?"

	if r.db != nil {
		rows, err = r.db.Query(cmd, gachaTypeID)
	} else {
		rows, err = r.tx.Query(cmd, gachaTypeID)
	}

	if err != nil {
		return gachaProbabilityGroups, xerrors.Errorf("Call Query: %w", err)
	}

	for rows.Next() {
		var gachaProbabilityGroupRow model.GachaProbabilityGroup
		if err := rows.Scan(&gachaProbabilityGroupRow.GachaTypeID, &gachaProbabilityGroupRow.Number, &gachaProbabilityGroupRow.GachaProbabilityGroupID, &gachaProbabilityGroupRow.Rate, &gachaProbabilityGroupRow.CreatedAt, &gachaProbabilityGroupRow.UpdatedAt); err != nil {
			return gachaProbabilityGroups, xerrors.Errorf("Call Scan: %w", err)
		}
		gachaProbabilityGroups = append(gachaProbabilityGroups, gachaProbabilityGroupRow)
	}

	if err := rows.Err(); err != nil {
		return gachaProbabilityGroups, xerrors.Errorf("Call rows.Err(): %w", err)
	}

	return gachaProbabilityGroups, nil
}
