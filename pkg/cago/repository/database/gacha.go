package database

import (
	"ca-tech-dojo-go/pkg/cago/model"
	"database/sql"
	"log"
)

type dbGachaRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func (r *dbGachaRepository) FindByGachaType(gachaTypeID int) ([]model.Gacha, error) {
	var gacha []model.Gacha
	var rows *sql.Rows
	var err error
	cmd := "SELECT id, chara_id, gacha_type_id, rate_type_id, rate, created_at, updated_at FROM gacha where gacha_type_id = ?"

	if r.db != nil {
		rows, err = r.db.Query(cmd, gachaTypeID)
	} else {
		rows, err = r.tx.Query(cmd, gachaTypeID)
	}

	if err != nil {
		log.Fatal(err)
		return gacha, err
	}

	for rows.Next() {
		var gachaRow model.Gacha
		if err := rows.Scan(&gachaRow.ID, &gachaRow.CharaID, &gachaRow.GachaTypeID, &gachaRow.RateTypeID, &gachaRow.Rate, &gachaRow.CreatedAt, &gachaRow.UpdatedAt); err != nil {
			log.Fatal(err)
			return gacha, err
		}
		gacha = append(gacha, gachaRow)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return gacha, err
	}

	return gacha, nil
}
