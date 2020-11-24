package database

import (
	"ca-tech-dojo-go/pkg/cago/model"
	"database/sql"
	"time"

	"golang.org/x/xerrors"
)

type dbUserRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func (r *dbUserRepository) Find(id int) (model.User, error) {
	var user model.User
	var row *sql.Row
	cmd := "SELECT id, name, token, created_at, updated_at FROM user WHERE id = ?"

	if r.db != nil {
		row = r.db.QueryRow(cmd, id)
	} else {
		row = r.tx.QueryRow(cmd, id)
	}
	if err := row.Scan(&user.ID, &user.Name, &user.Token, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return user, xerrors.Errorf("Call Scan: %w", err)
	}

	return user, nil
}

func (r *dbUserRepository) FindByToken(token string) (model.User, error) {
	var user model.User
	var row *sql.Row
	cmd := "SELECT id, name, token, created_at, updated_at FROM user WHERE token = ?"

	if r.db != nil {
		row = r.db.QueryRow(cmd, token)
	} else {
		row = r.tx.QueryRow(cmd, token)
	}
	if err := row.Scan(&user.ID, &user.Name, &user.Token, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return user, xerrors.Errorf("Call Scan: %w", err)
	}

	return user, nil
}

func (r *dbUserRepository) Create(user *model.User) error {
	tx := r.tx
	cmd := "INSERT INTO user (name, token, created_at) VALUES (?, ?, ?)"
	_, err := tx.Exec(cmd, user.Name, user.Token, user.CreatedAt)
	if err != nil {
		return xerrors.Errorf("Call Exec: %w", err)
	}
	return nil
}

func (r *dbUserRepository) UpdateName(name string, UpdatedAt time.Time, id int) error {
	tx := r.tx
	cmd := "UPDATE user SET name = ?, updated_at = ? WHERE id = ?"
	_, err := tx.Exec(cmd, name, UpdatedAt, id)
	if err != nil {
		return xerrors.Errorf("Call Exec: %w", err)
	}
	return nil
}

func (r *dbUserRepository) GetAllUserScore() ([]model.UserRanking, error) {
	var userRanks []model.UserRanking
	var rows *sql.Rows
	var err error

	cmd := "SELECT u.id, COUNT(cu.user_id) FROM user u LEFT OUTER JOIN chara_user cu ON u.id = cu.user_id GROUP BY u.id"

	if r.db != nil {
		rows, err = r.db.Query(cmd)
	} else {
		rows, err = r.tx.Query(cmd)
	}

	if err != nil {
		return userRanks, xerrors.Errorf("Call Query: %w", err)
	}

	for rows.Next() {
		var userRankRow model.UserRanking
		if err := rows.Scan(&userRankRow.UserID, &userRankRow.Score); err != nil {
			return userRanks, xerrors.Errorf("Call Scan: %w", err)
		}
		userRanks = append(userRanks, userRankRow)
	}

	if err := rows.Err(); err != nil {
		return userRanks, xerrors.Errorf("Call rows.Err(): %w", err)
	}

	return userRanks, nil
}
