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
	GachaProbabilityGroup() GachaProbabilityGroupQuery
}

type Transaction interface {
	User() UserCommand
	Chara() CharaCommand
}

type UserQuery interface {
	Find(id int) (model.User, error)
	FindByIDs(IDs []int) ([]model.User, error)
	FindByToken(token string) (model.User, error)
	GetAllUserScore() ([]model.UserRanking, error)
}

type UserCommand interface {
	UserQuery

	Create(user *model.User) error
	UpdateName(name string, UpdatedAt time.Time, id int) error
}

type GachaQuery interface {
	FindByGroupIDs(groupIDs []string) ([]model.GachaProbability, error)
}

type CharaQuery interface {
	FindByIDs(IDs []int) ([]model.Chara, error)
	FindUserCharaByUserID(UserID int) ([]model.CharaUser, error)
}

type CharaCommand interface {
	AddUserChara(charaIDs []int, CreatedAt time.Time, userID int) error
}

type GachaProbabilityGroupQuery interface {
	FindByGachaType(gachaTypeID int) ([]model.GachaProbabilityGroup, error)
}
