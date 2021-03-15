package service

import (
	"ca-tech-dojo-go/pkg/cago/model"
	"ca-tech-dojo-go/pkg/cago/repository"
	"ca-tech-dojo-go/pkg/cago/service/input"
	"ca-tech-dojo-go/pkg/cago/service/output"
	"context"
	"time"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

// UserService ユーザーサービス
type UserService interface {
	CreateUser(ctx context.Context, user *input.CreateUser) (output.CreateUser, error)
	GetUser(ctx context.Context, user *input.GetUser) (output.GetUser, error)
	UpdateUser(ctx context.Context, user *input.UpdateUser) (output.UpdateUser, error)
	GetUsers(ctx context.Context, users []*input.GetUser) ([]*output.GetUser, error)
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
		return outputUser, xerrors.Errorf("Call NewConnection: %w", err)
	}

	userModel, err := con.User().Find(user.ID)
	if err != nil {
		return outputUser, xerrors.Errorf("Call Find: %w", err)
	}

	outputUser.Name = userModel.Name
	return outputUser, nil
}

func (us *userService) GetUsers(ctx context.Context, users []*input.GetUser) ([]*output.GetUser, error) {
	var outputUsers []*output.GetUser

	con, err := us.r.NewConnection()
	if err != nil {
		return outputUsers, xerrors.Errorf("Call NewConnection: %w", err)
	}

	var IDs []int
	for _, user := range users {
		IDs = append(IDs, user.ID)
	}

	userModel, err := con.User().FindByIDs(IDs)
	if err != nil {
		return outputUsers, xerrors.Errorf("Call Find: %w", err)
	}

	for _, user := range userModel {
		outputUsers = append(outputUsers, &output.GetUser{
			ID:   user.ID,
			Name: user.Name,
		})
	}

	return outputUsers, nil
}

func (us *userService) CreateUser(ctx context.Context, user *input.CreateUser) (output.CreateUser, error) {
	var modelUser model.User
	uuidV4 := uuid.New()
	modelUser.Name = user.Name
	modelUser.Token = uuidV4.String()
	modelUser.CreatedAt = time.Now()

	var outputUser output.CreateUser

	con, err := us.r.NewConnection()
	if err != nil {
		return outputUser, xerrors.Errorf("Call NewConnection: %w", err)
	}

	err = con.RunTransaction(func(tx repository.Transaction) error {
		err := tx.User().Create(&modelUser)
		if err != nil {
			return xerrors.Errorf("Call Create: %w", err)
		}

		return nil
	})

	if err != nil {
		return outputUser, xerrors.Errorf("Call RunTransaction: %w", err)
	}

	outputUser.Xtoken = modelUser.Token
	return outputUser, nil
}

func (us *userService) UpdateUser(ctx context.Context, user *input.UpdateUser) (output.UpdateUser, error) {

	var outputUser output.UpdateUser

	con, err := us.r.NewConnection()
	if err != nil {
		return outputUser, xerrors.Errorf("Call NewConnection: %w", err)
	}

	err = con.RunTransaction(func(tx repository.Transaction) error {
		err := tx.User().UpdateName(user.Name, time.Now(), user.ID)
		if err != nil {
			return xerrors.Errorf("Call UpdateName: %w", err)
		}

		return nil
	})

	if err != nil {
		return outputUser, xerrors.Errorf("Call RunTransaction: %w", err)
	}

	return outputUser, nil
}
