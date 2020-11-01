package database

import (
	"ca-tech-dojo-go/pkg/cago/model"
	"database/sql"
	"log"
	"strings"
	"time"
)

type dbCharaRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func (r *dbCharaRepository) FindByIDs(IDs []int32) ([]model.Chara, error) {
	var chara []model.Chara
	var rows *sql.Rows
	var err error
	cmd := "SELECT id, name, created_at, updated_at FROM chara WHERE id IN (?" + strings.Repeat(",?", len(IDs)-1) + ")"

	var vars []interface{}

	for _, s := range IDs {
		vars = append(vars, s)
	}

	if r.db != nil {
		rows, err = r.db.Query(cmd, vars...)
	} else {
		rows, err = r.tx.Query(cmd, vars...)
	}

	if err != nil {
		log.Fatal(err)
		return chara, err
	}

	for rows.Next() {
		var gachaRow model.Chara
		if err := rows.Scan(&gachaRow.ID, &gachaRow.Name, &gachaRow.CreatedAt, &gachaRow.UpdatedAt); err != nil {
			log.Fatal(err)
			return chara, err
		}
		chara = append(chara, gachaRow)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return chara, err
	}

	return chara, nil
}

func (r *dbCharaRepository) AddUserChara(charaIDs []int32, CreatedAt time.Time, userID int32) error {
	tx := r.tx
	cmd := "INSERT INTO chara_user (user_id, chara_id, created_at) VALUES (?, ?, ?)" + strings.Repeat(",(?, ?, ?)", len(charaIDs)-1)

	var vars []interface{}

	for _, s := range charaIDs {
		vars = append(vars, userID)
		vars = append(vars, s)
		vars = append(vars, CreatedAt)
	}
	_, err := tx.Exec(cmd, vars...)

	return err
}
