package database

import (
	"ca-tech-dojo-go/pkg/cago/model"
	"database/sql"
	"strings"
	"time"

	"golang.org/x/xerrors"
)

type dbCharaRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func (r *dbCharaRepository) FindByIDs(IDs []int) ([]model.Chara, error) {
	var chara []model.Chara
	var rows *sql.Rows
	var err error

	if len(IDs) <= 0 {
		return chara, nil
	}

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
		return chara, xerrors.Errorf("Call Query: %w", err)
	}

	for rows.Next() {
		var charaRow model.Chara
		if err := rows.Scan(&charaRow.ID, &charaRow.Name, &charaRow.CreatedAt, &charaRow.UpdatedAt); err != nil {
			return chara, xerrors.Errorf("Call Scan: %w", err)
		}
		chara = append(chara, charaRow)
	}

	if err := rows.Err(); err != nil {
		return chara, xerrors.Errorf("Call rows.Err(): %w", err)
	}

	return chara, nil
}

func (r *dbCharaRepository) FindUserCharaByUserID(UserID int) ([]model.CharaUser, error) {
	var userCharas []model.CharaUser
	var rows *sql.Rows
	var err error
	cmd := "SELECT id, user_id, chara_id, created_at, updated_at FROM chara_user WHERE user_id = ?"

	if r.db != nil {
		rows, err = r.db.Query(cmd, UserID)
	} else {
		rows, err = r.tx.Query(cmd, UserID)
	}

	if err != nil {
		return userCharas, xerrors.Errorf("Call Query: %w", err)
	}

	for rows.Next() {
		var charaUserRow model.CharaUser
		if err := rows.Scan(&charaUserRow.ID, &charaUserRow.UserID, &charaUserRow.CharaID, &charaUserRow.CreatedAt, &charaUserRow.UpdatedAt); err != nil {
			return userCharas, xerrors.Errorf("Call Scan: %w", err)
		}
		userCharas = append(userCharas, charaUserRow)
	}

	if err := rows.Err(); err != nil {
		return userCharas, xerrors.Errorf("Call rows.Err(): %w", err)
	}

	return userCharas, nil
}

func (r *dbCharaRepository) AddUserChara(charaIDs []int, CreatedAt time.Time, userID int) error {

	if len(charaIDs) <= 0 {
		return nil
	}

	tx := r.tx
	cmd := "INSERT INTO chara_user (user_id, chara_id, created_at) VALUES (?, ?, ?)" + strings.Repeat(",(?, ?, ?)", len(charaIDs)-1)

	vars := make([]interface{}, 0, len(charaIDs)*3)

	for _, charaID := range charaIDs {
		vars = append(vars, userID, charaID, CreatedAt)
	}
	_, err := tx.Exec(cmd, vars...)

	if err != nil {
		return xerrors.Errorf("Call Exec: %w", err)
	}
	return nil
}
