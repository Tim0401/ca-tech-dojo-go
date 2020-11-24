package database

import (
	"ca-tech-dojo-go/pkg/cago/model"
	"database/sql"
	"log"
	"strings"

	"golang.org/x/xerrors"
)

type dbGachaRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func (r *dbGachaRepository) FindByGroupIDs(groupIDs []string) ([]model.GachaProbability, error) {
	var gacha []model.GachaProbability
	var rows *sql.Rows
	var err error

	if len(groupIDs) <= 0 {
		return gacha, nil
	}

	cmd := "SELECT group_id, number, chara_id, rate, created_at, updated_at FROM gacha_probability WHERE group_id IN (?" + strings.Repeat(",?", len(groupIDs)-1) + ")"

	vars := make([]interface{}, 0, len(groupIDs))
	for _, id := range groupIDs {
		vars = append(vars, id)
	}

	if r.db != nil {
		rows, err = r.db.Query(cmd, vars...)
	} else {
		rows, err = r.tx.Query(cmd, vars...)
	}

	if err != nil {
		return gacha, xerrors.Errorf("Call Query: %w", err)
	}

	for rows.Next() {
		var gachaRow model.GachaProbability
		if err := rows.Scan(&gachaRow.GroupID, &gachaRow.Number, &gachaRow.CharaID, &gachaRow.Rate, &gachaRow.CreatedAt, &gachaRow.UpdatedAt); err != nil {
			log.Fatal(err)
			return gacha, xerrors.Errorf("Call Scan: %w", err)
		}
		gacha = append(gacha, gachaRow)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return gacha, xerrors.Errorf("Call rows.Err(): %w", err)
	}

	return gacha, nil
}
