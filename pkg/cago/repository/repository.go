package repository

import (
	"ca-tech-dojo-go/pkg/cago/model"
)

type Repository interface {
	NewConnection() (Connection, error)
	MustConnection() Connection
}

type Connection interface {
	Close() error
	RunTransaction(func(tx Transaction) error) error

	User() UserQuery
}

type Transaction interface {
	User() UserCommand
}

type UserQuery interface {
	Find(id int) (*model.User, error)
	FindByToken(token string) (*model.User, error)
}

type UserCommand interface {
	UserQuery

	Create(user *model.User) error
}
