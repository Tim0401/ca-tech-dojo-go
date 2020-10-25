package database

import (
	"ca-tech-dojo-go/pkg/cago/model"
	"database/sql"
	"log"
)

type dbUserRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func (r *dbUserRepository) Find(id int) (*model.User, error) {
	var user model.User
	var row *sql.Row
	cmd := "SELECT id, name, token, created_at, updated_at FROM user WHERE id = ?"

	if r.db != nil {
		row = r.db.QueryRow(cmd, id)
	} else {
		row = r.db.QueryRow(cmd, id)
	}
	if err := row.Scan(&user.ID, &user.Name, &user.Token, &user.CreatedAt, &user.UpdatedAt); err != nil {
		log.Println(err)
		return nil, err
	}

	return &user, nil
}

func (r *dbUserRepository) FindByToken(token string) (*model.User, error) {
	var user model.User
	var row *sql.Row
	cmd := "SELECT id, name, token, created_at, updated_at FROM user WHERE token = ?"

	if r.db != nil {
		row = r.db.QueryRow(cmd, token)
	} else {
		row = r.db.QueryRow(cmd, token)
	}
	if err := row.Scan(&user.ID, &user.Name, &user.Token, &user.CreatedAt, &user.UpdatedAt); err != nil {
		log.Println(err)
		return nil, err
	}

	return &user, nil
}

func (r *dbUserRepository) Create(user *model.User) error {
	tx := r.tx
	cmd := "INSERT INTO user (name, token, created_at) VALUES (?, ?, ?)"
	_, err := tx.Exec(cmd, user.Name, user.Token, user.CreatedAt)
	return err
}

func (r *dbUserRepository) UpdateNameByToken(name string, token string) error {
	tx := r.tx
	cmd := "UPDATE user SET name = ? WHERE token = ?"
	_, err := tx.Exec(cmd, name, token)
	return err
}
