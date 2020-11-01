package repository

import (
	"ca-tech-dojo-go/pkg/cago/model"
	"time"
)

type Repository interface {
	NewConnection() (Connection, error)
	MustConnection() Connection
}

type Connection interface {
	Close() error
	RunTransaction(func(tx Transaction) error) error

	User() UserQuery
	Gacha() GachaQuery
}

type Transaction interface {
	User() UserCommand
}

type UserQuery interface {
	Find(id int32) (model.User, error)
	FindByToken(token string) (model.User, error)
}

type UserCommand interface {
	UserQuery

	Create(user *model.User) error
	UpdateName(name string, UpdatedAt time.Time, id int32) error
}

type GachaQuery interface {
	FindAll() ([]model.Gacha, error)
}
