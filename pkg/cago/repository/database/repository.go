package database

import (
	"ca-tech-dojo-go/pkg/cago/repository"
	"database/sql"
)

// NewRepository generates a new repository using DB.
func NewRepository(db *sql.DB) repository.Repository {
	return &dbRepository{
		db: db,
	}
}

type dbRepository struct {
	db *sql.DB
}

type dbConnection struct {
	db *sql.DB
}

type dbTransaction struct {
	db *sql.Tx
}

func (r *dbRepository) NewConnection() (repository.Connection, error) {
	return &dbConnection{
		db: r.db,
	}, nil
}

func (r *dbRepository) MustConnection() repository.Connection {
	con, err := r.NewConnection()
	if err != nil {
		panic(err)
	}

	return con
}

func (con *dbConnection) Close() error {
	return nil
}

func (con *dbConnection) RunTransaction(f func(repository.Transaction) error) error {
	tx, _ := con.db.Begin()

	err := f(&dbTransaction{db: tx})
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (con *dbConnection) User() repository.UserQuery {
	return &dbUserRepository{db: con.db}
}

func (tx *dbTransaction) User() repository.UserCommand {
	return &dbUserRepository{tx: tx.db}
}

func (con *dbConnection) Gacha() repository.GachaQuery {
	return &dbGachaRepository{db: con.db}
}
