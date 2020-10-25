package service

import (
	"ca-tech-dojo-go/pkg/cago/model"
	"ca-tech-dojo-go/pkg/cago/repository"
	"ca-tech-dojo-go/pkg/cago/service/input"
	"ca-tech-dojo-go/pkg/cago/service/output"
	"context"
	"time"

	"github.com/google/uuid"
)

// UserService ユーザーサービス
type UserService interface {
	CreateUser(ctx context.Context, user *input.CreateUser) (output.CreateUser, error)
	GetUser(ctx context.Context, user *input.GetUser) (output.GetUser, error)
	UpdateUser(ctx context.Context, user *input.UpdateUser) (output.UpdateUser, error)
}

type userService struct {
	r repository.Repository
}

// NewUserService ユーザーサービス作成
func NewUserService(r repository.Repository) UserService {
	return &userService{r}
}

func (us *userService) GetUser(ctx context.Context, user *input.GetUser) (output.GetUser, error) {
	var outputUser output.GetUser

	con, err := us.r.NewConnection()
	if err != nil {
		return outputUser, err
	}
	defer con.Close()

	userModel, err := con.User().FindByToken(user.Xtoken)

	outputUser.Name = userModel.Name
	return outputUser, err
}

func (us *userService) CreateUser(ctx context.Context, user *input.CreateUser) (output.CreateUser, error) {
	var modelUser model.User
	uuidV4 := uuid.New()
	modelUser.Name = user.Name
	modelUser.Token = uuidV4.String()
	modelUser.CreatedAt = time.Now()

	// 保存
	// us.ur.RunTransaction(func(tx *sql.Tx) error {
	// 	err := us.ur.CreateUser(ctx, &modelUser, tx)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	return nil
	// })

	var outputUser output.CreateUser

	con, err := us.r.NewConnection()
	if err != nil {
		return outputUser, err
	}
	defer con.Close()

	err = con.RunTransaction(func(tx repository.Transaction) error {
		err := tx.User().Create(&modelUser)
		if err != nil {
			return err
		}

		return nil
	})

	outputUser.Xtoken = modelUser.Token
	return outputUser, err
}

func (us *userService) UpdateUser(ctx context.Context, user *input.UpdateUser) (output.UpdateUser, error) {

	var outputUser output.UpdateUser

	con, err := us.r.NewConnection()
	if err != nil {
		return outputUser, err
	}
	defer con.Close()

	err = con.RunTransaction(func(tx repository.Transaction) error {
		err := tx.User().UpdateNameByToken(user.Name, time.Now(), user.Xtoken)
		if err != nil {
			return err
		}

		return nil
	})

	return outputUser, err
}
