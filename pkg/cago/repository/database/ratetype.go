package database

import (
	"ca-tech-dojo-go/pkg/cago/model"
	"database/sql"
	"log"
)

type dbRateTypeRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func (r *dbRateTypeRepository) FindByGachaType(gachaTypeID int) ([]model.RateType, error) {
	var rateType []model.RateType
	var rows *sql.Rows
	var err error
	cmd := "SELECT r.id, r.name, r.rate, r.created_at, r.updated_at FROM rate_type r " +
		"INNER JOIN gacha_type_rate_type gr ON r.id = gr.rate_type_id where gr.gacha_type_id = ?"

	if r.db != nil {
		rows, err = r.db.Query(cmd, gachaTypeID)
	} else {
		rows, err = r.tx.Query(cmd, gachaTypeID)
	}

	if err != nil {
		log.Fatal(err)
		return rateType, err
	}

	for rows.Next() {
		var rateTypeRow model.RateType
		if err := rows.Scan(&rateTypeRow.ID, &rateTypeRow.Name, &rateTypeRow.Rate, &rateTypeRow.CreatedAt, &rateTypeRow.UpdatedAt); err != nil {
			log.Fatal(err)
			return rateType, err
		}
		rateType = append(rateType, rateTypeRow)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return rateType, err
	}

	return rateType, nil
}
