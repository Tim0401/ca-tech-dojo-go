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
	RateType() RateTypeQuery
}

type Transaction interface {
	User() UserCommand
	Chara() CharaCommand
}

type UserQuery interface {
	Find(id int) (model.User, error)
	FindByToken(token string) (model.User, error)
}

type UserCommand interface {
	UserQuery

	Create(user *model.User) error
	UpdateName(name string, UpdatedAt time.Time, id int) error
}

type GachaQuery interface {
	FindByGachaType(gachaTypeID int) ([]model.Gacha, error)
}

type CharaQuery interface {
	FindByIDs(IDs []int) ([]model.Chara, error)
	FindUserCharaByUserID(UserID int) ([]model.CharaUser, error)
}

type CharaCommand interface {
	AddUserChara(charaIDs []int, CreatedAt time.Time, userID int) error
}

type RateTypeQuery interface {
	FindByGachaType(gachaTypeID int) ([]model.RateType, error)
}
