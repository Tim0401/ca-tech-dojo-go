package database

import (
	"ca-tech-dojo-go/pkg/cago/model"
	"database/sql"

	"golang.org/x/xerrors"
)

type dbCharaMemoryRepository struct {
	db *sql.DB
	tx *sql.Tx
}

var charaMap = make(map[int]model.Chara)

func InitChara(db *sql.DB) error {
	cmd := "SELECT id, name, created_at, updated_at FROM chara"
	rows, err := db.Query(cmd)
	if err != nil {
		return err
	}

	for rows.Next() {
		var charaRow model.Chara
		if err := rows.Scan(&charaRow.ID, &charaRow.Name, &charaRow.CreatedAt, &charaRow.UpdatedAt); err != nil {
			return xerrors.Errorf("Call Scan: %w", err)
		}
		charaMap[charaRow.ID] = charaRow
	}

	if err := rows.Err(); err != nil {
		return xerrors.Errorf("Call rows.Err(): %w", err)
	}
	return nil
}

func (r *dbCharaMemoryRepository) FindByIDs(IDs []int) ([]model.Chara, error) {
	var chara []model.Chara

	for _, id := range IDs {
		charaRow, ok := charaMap[id]
		if ok {
			chara = append(chara, charaRow)
		}
	}
	return chara, nil
}

func (r *dbCharaMemoryRepository) FindUserCharaByUserID(UserID int) ([]model.CharaUser, error) {
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
