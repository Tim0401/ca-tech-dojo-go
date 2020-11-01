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
	Chara() CharaQuery
}

type Transaction interface {
	User() UserCommand
	Chara() CharaCommand
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

type CharaQuery interface {
	FindByIDs(IDs []int32) ([]model.Chara, error)
}

type CharaCommand interface {
	AddUserChara(charaIDs []int32, CreatedAt time.Time, userID int32) error
}
