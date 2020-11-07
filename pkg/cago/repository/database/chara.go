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

func (r *dbCharaRepository) FindByIDs(IDs []int) ([]model.Chara, error) {
	var chara []model.Chara
	var rows *sql.Rows
	var err error
	cmd := "SELECT id, name, created_at, updated_at FROM chara WHERE id IN (?" + strings.Repeat(",?", len(IDs)-1) + ")"

	vars := make([]interface{}, 0, len(IDs))

	for _, id := range IDs {
		vars = append(vars, id)
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

func (r *dbCharaRepository) AddUserChara(charaIDs []int, CreatedAt time.Time, userID int) error {
	tx := r.tx
	cmd := "INSERT INTO chara_user (user_id, chara_id, created_at) VALUES (?, ?, ?)" + strings.Repeat(",(?, ?, ?)", len(charaIDs)-1)

	vars := make([]interface{}, 0, len(charaIDs)*3)

	for _, charaID := range charaIDs {
		vars = append(vars, userID, charaID, CreatedAt)
	}
	_, err := tx.Exec(cmd, vars...)

	return err
}
